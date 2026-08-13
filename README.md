# CA Prop. 65 Daily Exposure Calculator

This small application determines the daily exposure to a cosmetic product containing a Prop. 65 chemical. The result can then be compared to the corresponding Safe Harbor Level to determine if a Prop. 65 warning is required for the product.

## Calculation

The following details the step by step mathematical model used to determine the daily exposure to a residual CA Prop. 65 chemical within an intentionally added material. The resulting residual chemical exposure in the cosmetic formulation can be compared against California Proposition 65 Safe Harbor Levels (NSRL/MADL) to determine if a warning is required.

### 1. Chemical concentration in finished product

This formula calculates the concentration of the residual chemical in the finished cosmetic product based on the raw material inclusion rate:

```math
C_{\text{chem, finished}} = C_{\text{chem, raw}} \times P_{\text{raw}}
```
**_where_**
- $`C_{\text{chem, finished}}`$ ≡ concentration of the Prop. 65 chemical in the finished product, in units of μg/g or ppm
- $`C_{\text{chem, raw}}`$ ≡ maximum residual contaminant level in the raw ingredient, in units of μg/g or ppm
  - We use maximum contaminant level because it is best to use worst case scenario for safety
- $`P_{\text{raw}}`$ ≡ inclusion fraction of raw ingredient in final formula
  - i.e., W/W%
  - e.g., 14% = 0.14 

### 2. Effective Daily Applied Mass

This formula calculates the mass of the product remaining on skin/hair after normal consumer usage:

```math
m_{\text{effective}} = m_{\text{usage, avg}} \times R
```
**_where_**
- $`m_{\text{effective}}`$ ≡ the mass of product that remains after normal usage
- $`m_{\text{usage, avg}}`$ ≡ the average mass of finished product used.
  - These values can be viewed below in sources
- $`R`$ ≡ product retention factor, unitless
  - These values can be found below in sources 
  - Typically 1.0 for leave-in and 0.1 or 0.01 for rinse-off

### 3. Total Daily Chemical Exposure

This formula calculates the total daily mass of the residual chemical absorbed/exposed per day for typical usage:

```math
E_{\text{daily}} = C_{\text{chem, finished}} \times m_{\text{usage, avg}} \times R \times AF
```

_or simplified using_ $`m_{\text{effective}}`$:

```math
E_{\text{daily}} = C_{\text{chem, finished}} \times m_{\text{effective}} \times AF
```
**_where_**
- $`E_{\text{daily}}`$ ≡ the total daily exposure to the Prop. 65 chemical, in units of μg/day
- $`C_{\text{chem, finished}}`$ ≡ concentration of the Prop. 65 chemical in the finished product, in units of μg/g 
- $`m_{\text{usage, avg}}`$ ≡ the average mass of finished product used.
- $`R`$ ≡ product retention factor, unitless
- $`AF`$ ≡ dermal absorption factor, unitless
  - This defaults to a conservative 1.00, as anything below 1.00 would 
  - require empirical toxicological dermal penetration data

### 4. Prop. 65 Safe Harbor Evaluation

```math
\text{Status} = \begin{cases} \text{Exempt (No Warning Required)} & \text{if } E_{\text{daily}} \le \text{Safe Harbor Level} \\ \text{Warning Required} & \text{if } E_{\text{daily}} > \text{Safe Harbor Level} \end{cases}
```
- Safe Harbor Levels can be found on the [OEHHA Webiste](https://oehha.ca.gov/proposition-65/chemicals-considered-or-listed-under-proposition-65)
- This is an example for [Dichloroacetic acid](https://oehha.ca.gov/proposition-65/chemicals/dichloroacetic-acid)

## Data Sources

- [Environmental Protection Agency's Exposure Factors Handbook Chapter 17](https://www.epa.gov/sites/default/files/2015-09/documents/efh-chapter17.pdf)
  - Table 17-41. Average Amount of Product Applied per Applicationa (grams) 
- [THE SCCS NOTES OF GUIDANCE FOR THE TESTING OF COSMETIC INGREDIENTS AND THEIR SAFETY EVALUATION 12TH REVISION](https://health.ec.europa.eu/document/download/32a999f7-d820-496a-b659-d8c296cc99c1_en?filename=sccs_o_273_final.pdf)
  - Table 3A (Page 27)
- [OEHHA Chemicals Considered or Listed Under Prop. 65](https://oehha.ca.gov/proposition-65/chemicals-considered-or-listed-under-proposition-65)

## Disclaimer
**USE AT YOUR OWN RISK.**

This tool is provided as-is, and the developer(s) are not liable for any errors, inaccuracies, or issues that may arise from its use. It is essential to exercise your own due diligence and verify the results obtained through this tool. The developer(s) make no guarantees about the accuracy, completeness, or reliability of the calculated data.

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
