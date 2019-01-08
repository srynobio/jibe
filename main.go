package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	arg "github.com/alexflint/go-arg"
	"github.com/brentp/vcfgo"
	"github.com/brentp/xopen"
)

// Current jibe version
const Version = "1.0.2"

func digest(bv []byte) string {
	hasher := sha1.New()
	hasher.Write(bv)

	bs := hasher.Sum(nil)
	sha := base64.StdEncoding.EncodeToString(bs)
	return sha
}

func echeck(m string, e error) {
	if e != nil {
		log.Println(m)
		panic(e)
	}
}

func main() {
	var args struct {
		VCF      []string `help:"VCF file to collect concordance from. Space separated."`
		SNP      bool     `help:"Only consider SNP calls."`
		NOMULTI  bool     `help:"Confirm via exit, no multi-allelic variants."`
		DATALINE bool     `help:"Will use complete variant dataline including INFO and (single|multi) Genotype fields"`
		CPUS     int      `help:"Number of CPUS workers to allow."`
		VERSION  bool     `help:"Print current version and exit."`
	}
	args.CPUS = runtime.NumCPU()
	arg.MustParse(&args)

	// Version print
	if args.VERSION == true {
		fmt.Println(Version)
		os.Exit(0)
	}

	// Check for two files.
	if len(args.VCF) != 2 {
		log.Panic("Two VCF files required.")
	}

	// Create wg and result channel.
	var wg sync.WaitGroup
	dchan := make(chan string)

	// start the workers.
	for _, file := range args.VCF {
		wg.Add(1)

		log.Println("Processing started for file: ", file)

		go func(f string) {
			defer wg.Done()
			// open VCF file.
			oFile, err := xopen.Ropen(f)
			echeck("Can't open VCF file.", err)
			defer oFile.Close()

			r, err := vcfgo.NewReader(oFile, false)
			echeck("Can't access vcf file.", err)
			defer r.Close()

			for {
				read := r.Read()
				if read == nil {
					break
				}

				// if NoMulti is requested.
				if args.NOMULTI == true {
					if len(read.Alt()) > 1 {
						log.Println("Multi-allelic variant found.")
						log.Println(read.String())
						os.Exit(1)
					}
				}

				// if SNP option is requested and skip if found.
				var altLength, refLength int
				if args.SNP == true {
					refLength = len(read.Reference)
					for _, dna := range read.Alt() {
						altLength = len(dna)
					}
				}
				if args.SNP == true && altLength > 1 || refLength > 1 {
					continue
				}

				// Start work.
				wg.Add(1)
				go func(record *vcfgo.Variant) {
					defer wg.Done()
					var site string
					if args.DATALINE == true {
						site = record.String()
					} else {
						site = fmt.Sprintf("%s:%d:%d:%s:%s", record.Chromosome, record.Start(), record.End(), record.Reference, record.Alt())
					}
					siteDigest := digest([]byte(site))
					dchan <- siteDigest
				}(read)
			}
		}(file)
	}

	// Close chan after all work completed.
	go func() {
		wg.Wait()
		defer close(dchan)
	}()

	// Get value from channel and make lookup.
	var union float64
	dataLookup := make(map[string]int)
	for receive := range dchan {
		if _, ok := dataLookup[receive]; ok {
			union++
			delete(dataLookup, receive)
			continue
		} else {
			dataLookup[receive]++
		}
	}
	// Make uniq a float as well.
	uniq := float64(len(dataLookup))

	// Run some checks.
	if union < 1 && len(dataLookup) < 1 {
		log.Panicln("No data found to review, please check file.")
	}
	if len(dataLookup) == 0.00 {
		fmt.Println("--- Results ---")
		fmt.Println("File Pair in 100% union.")
		fmt.Println("---------------")
		os.Exit(0)
	}

	// calculate totals
	countTotal := union + uniq
	precentUnion := (union / countTotal) * 100
	precentUniq := (uniq / countTotal) * 100

	// Simple Report.
	fmt.Println("--- Results ---")
	fmt.Printf("Precent union: %.2f%%\n", precentUnion)
	fmt.Printf("Precent uniq: %.2f%%\n", precentUniq)
	fmt.Println("---------------")
}
