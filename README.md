![test](https://github.com/srynobio/jibe/blob/dev/image/jibelogo.png)

### Introduction

Simple validation tool used to compare two versions of a VCF file.  Designed and beneficial to test the consistancy of called variants across different callers and pipelines at different time points.

It uses a simple hashing scheme which will compares SNP and INDEL calls without need for normalization first.  Header information will not be considered.

By default only the following will be compared:

`chromosome:start:end:reference:alt`

### Usage:

```
Usage: jibe [--vcf VCF] [--snp] [--nomulti] [--dataline] [--cpus CPUS]

Options:
  --vcf VCF              VCF file to collect concordance from. Space separated.
  --snp                  Only consider SNP calls.
  --nomulti              Confirm via exit, no multi-allelic variants.
  --dataline             Will use complete variant dataline including INFO and (single|multi) Genotype fields
  --cpus CPUS            Number of CPUS workers to allow. [default: NumCPU]
  --help, -h             display this help and exit
```
