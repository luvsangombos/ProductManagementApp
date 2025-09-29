# ProductManagementApp

A CLI application built with Go for managing and displaying product information in multiple formats (JSON, XML, CSV).

## Features

- **Product Management**: Store and manage product information including ID, name, date supplied, quantity, and unit price
- **Multiple Output Formats**: Export product data in JSON, XML, or CSV format
- **Smart Sorting**: Automatically sorts products by name (ascending) and then by unit price (descending)
- **CLI Interface**: Simple command-line interface for easy usage

## Product Structure

Each product contains the following fields:
- `ProductId`: Unique identifier (string)
- `Name`: Product name (string)
- `DateSupplied`: Date when product was supplied (time.Time)
- `Quantity`: Available quantity (int)
- `UnitPrice`: Price per unit (float64)

## Installation

1. Clone the repository:
```bash
git clone git@github.com:luvsangombos/ProductManagementApp.git
cd ProductManagementApp
```

2. Initialize Go module (if not already done):
```bash
go mod init ProductManagementApp
```

## Usage

Run the application with the desired output format:

```bash
# JSON format
go run main.go json

# XML format
go run main.go xml

# CSV format
go run main.go csv
```

### Sample Output

**JSON Format:**
```json
{
  "products": [
    {
      "ProductId": "29274582650152771644",
      "Name": "Apple",
      "DateSupplied": "2024-12-09T00:00:00Z",
      "Quantity": 18,
      "UnitPrice": 1.09
    }
  ]
}
```

**XML Format:**
```xml
<ProductList>
  <products>
    <Product>
      <ProductId>29274582650152771644</ProductId>
      <Name>Apple</Name>
      <DateSupplied>2024-12-09T00:00:00Z</DateSupplied>
      <Quantity>18</Quantity>
      <UnitPrice>1.09</UnitPrice>
    </Product>
  </products>
</ProductList>
```

**CSV Format:**
```
ProductId,Name,DateSupplied,Quantity,UnitPrice
29274582650152771644,Apple,2024-12-09,18,1.09
31288741190182539913,Banana,2025-02-13,240,0.65
31288741190182539912,Banana,2025-01-24,124,0.55
91899274600128155167,Carrot,2025-03-31,89,2.99
```

## Sample Data

The application comes with pre-loaded sample data:
- 2 Banana products with different prices and dates
- 1 Apple product
- 1 Carrot product

## Sorting Logic

Products are automatically sorted by:
1. **Name** (ascending alphabetical order)
2. **Unit Price** (descending order for products with the same name)

## Error Handling

The application includes proper error handling for:
- Invalid command line arguments
- JSON/XML marshaling errors
- Missing required parameters

## Requirements

- Go 1.16 or higher
- Git (for cloning)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Author

Created as part of Applied Software Development coursework.
