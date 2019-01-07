![jibe](https://github.com/srynobio/jibe/blob/dev/image/jibelogo.png)

### Introduction

* Simple validation tool used to compare two versions of a VCF file.  Designed and beneficial to test the consistency of called variants across different callers and pipelines at different time points.

* It uses a simple hashing scheme which will compares SNP and INDEL calls without need for normalization first.  Header information will not be considered.

* Runtime for WES data is quite quick, but will increase with WGS based on file size and number of individuals within.

* By default only the following will be compared:

`chromosome:start:end:reference:alt`

### Usage:

```
Usage: jibe [--vcf VCF] [--snp] [--nomulti] [--dataline] [--cpus CPUS] [--version]

Options:
  --vcf VCF              VCF file to collect concordance from. Space separated.
  --snp                  Only consider SNP calls.
  --nomulti              Confirm via exit, no multi-allelic variants.
  --dataline             Will use complete variant dataline including INFO and (single|multi) Genotype fields
  --cpus CPUS            Number of CPUS workers to allow. [default: NumCPU]
  --version              Print current version and exit.
  --help, -h             display this help and exit
```

### Instalation:

Prebuilt binaries for mac OS and Linux are available [here](https://github.com/srynobio/jibe/releases)
