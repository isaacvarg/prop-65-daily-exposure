# CA Prop. 65 Daily Exposure Calculator

This small application determines the daily exposure to a cosmetic product containing a Prop. 65 chemical. The result can then be compared to the corresponding Safe Harbor Level to determine if a Prop. 65 warning is required for the product.

## Disclaimer
**USE AT YOUR OWN RISK.**

This tool is provided as-is, and the developer(s) are not liable for any errors, inaccuracies, or issues that may arise from its use. It is essential to exercise your own due diligence and verify the results obtained through this tool. The developer(s) make no guarantees about the accuracy, completeness, or reliability of the calculated data.

## Calculation

An explanation and rationale for the calculation methods will be provided here in the future.


## Data Sources

- [Environmental Protection Agency's Exposure Factors Handbook Chapter 17](https://www.epa.gov/sites/default/files/2015-09/documents/efh-chapter17.pdf)
  - Table 17-41. Average Amount of Product Applied per Applicationa (grams) 
- [THE SCCS NOTES OF GUIDANCE FOR THE TESTING OF COSMETIC INGREDIENTS AND THEIR SAFETY EVALUATION 12TH REVISION](https://health.ec.europa.eu/document/download/32a999f7-d820-496a-b659-d8c296cc99c1_en?filename=sccs_o_273_final.pdf)
  - Table 3A (Page 27)


## Installation

Packaged binaries and installation methods are coming. 

For now you can use this by cloning the repo and building:
```bash
# clone and build
git clone https://github.com/isaacvarg/prop-65-daily-exposure.git
cd prop-65-daily-exposure
go build .

# run
./prop-65-daily-exposure
```
Also you could copy this to your user binaries if it is in your path:
```bash
# build the go application similar to above
cp prop-65-daily-exposure ~/.local/bin/

# now you can run anywhere by
prop-65-daily-exposure
```

## What is California Proposition 65

Proposition 65, officially known as the Safe Drinking Water and Toxic Enforcement Act of 1986, was enacted as a ballot initiative in November 1986. The law protects the state's drinking water sources from being contaminated with chemicals known to cause cancer, birth defects or other reproductive harm, and requires businesses to inform Californians about exposures to such chemicals.

Proposition 65 requires the state to maintain and update the list of chemicals subject to the law’s requirements.
