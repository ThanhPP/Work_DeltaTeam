# Forward & Short link tool

_________

## Môi trường cài đặt

1. **Ngôn ngữ** : Golang
2. **Quản lý package** : Go modules

_________

## Cài đặt

_ Tạo 1 folder config trong đó có chứa file config.go
- Đổi đuôi file config_template.txt -> config.go.
- Các đường dẫn có thể thay đổi tùy ý.
- Đăng ký các API Key : name.com và rebrandly.com.

```go
package config

//NameAPIKey1 API key for name.com
var NameAPIKey1 = ""

//NameAPIKey2 API key for name.com (Concurrency)
var NameAPIKey2 = ""

//NameUsername Username for name.com
var NameUsername = ""

//NameDomain customize domain for name.com
var NameDomain = ""

//StoreLinkPath Paste store link here
var StoreLinkPath = "...001StoreLink.txt"

//TempForwardLinkPath Paste temp forward link here
var TempForwardLinkPath = "...002ForwardLink.txt"

//TempResultForwardLinkPath store Forward link created
var TempResultForwardLinkPath = "...003TempResultForwardLink.txt"

//RebrandlyAPIKey APIKey use for rebrandly
var RebrandlyAPIKey = ""

//RebrandlyDomainID Domain ID for rebrandly
var RebrandlyDomainID = ""

//SlashTagPath store the slashtag use for rebrandly.com
var SlashTagPath = "...004Slashtag.txt"

```

_ Chạy các câu lệnh để cài đặt :

```bash
go mod init
go build -v
```

_ Sử dụng :

```bash
./Work_ShortLinkTool -forward=false -shortLink=false
```

_________

## Sử dụng

_ Có các flag để lựa chọn việc sử dụng

- forward : boolean
- shortLink : boolean

_ Các flag này đều có giá trị mặc định là false, khi chạy chương trình cần thêm các tham số tại cmd để chuyển sang true.

_ forward : true

- Lấy dữ liệu store link tại : *config.StoreLinkPath*
- Lấy dữ liệu forwardlink (không bao gồm domain name) tại : *config.TempForwardLinkPath*
- Lưu các forward link được tạo vào : *config.TempResultForwardLinkPath* (Sử dụng cho việc short link)

_ shortLink : true

- Lấy dữ liệu forward link tại : *config.TempResultForwardLinkPath*
- Lấy dữ liệu slashtag tại : *config.SlashTagPath*

_________

## Lưu ý

_ Các đầu vào và đầu ra nên có định dạng .txt

_ Chương trình sẽ kiểm tra số lượng link của các file và chỉ chạy khi số lượng bằng nhau

- *config.StoreLinkPath* : Khi dùng forward
- *config.TempResultForwardLinkPath* : Khi dùng forward
- *config.TempResultForwardLinkPath* : Khi dùng short link
- *config.SlashTagPath* : Khi dùng short link
