# Dev deva test

1. ให้ทำ func สำหรับ gen ข้อมูล 100 rows ค่ะ โดยค่าที่ gen จะอยู่ที่ระหว่าง 1-1000 มี 2 ค่า active_power กับ input_power

2. เมื่อได้ข้อมูลมาจากข้อ 1 ให้ทำ query param เพื่อ sum ผมรวมของค่า active_power กับ input_power ทั้งหมด

3. จากนั้นให้ response ผลลัพท์เป็น JSON

## Installation

1. Make sure you have Go (Golang) installed. You can download it from [here](https://golang.org/dl/).

2. Install project dependencies:

   ```sh
   go mod tidy
   ```

## Usage

1. Run the application:

   ```sh
   go run cmd/main.go
   ```

2. Access the API endpoints using a web browser or an API client.

## Endpoints

- `?query=active_power` Calculates and returns the sum of `active_power` for generated data.
- `?query=input_power` Calculates and returns the sum of `input_power` for generated data.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
