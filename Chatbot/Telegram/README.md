# TELEGRAM CHATBOT

## Giới thiệu chung :

_ Chương trình sẽ tạo ra 1 chatbot trên telegram để phục vụ việc tạo forward links và short link.

_ Đầu vào là range trong google sheet, chương trình sẽ sử dụng API của google để lấy data và chạy.

## Tài liệu tham khảo thêm :

https://www.prudentdevs.club/gsheets-go

https://developers.google.com/sheets/api/quickstart/go#step_3_set_up_the_sample

https://github.com/THANHPP/Work_DeltaTeam/Work_ShortLinkTool

## Thư viện sử dụng :

_ Telegram chatbot : github.com/go-telegram-bot-api/telegram-bot-api

_ Google sheet : 

- "golang.org/x/oauth2/google"

- "google.golang.org/api/sheets/v4"


_ ENV file : 

- github.com/joho/godotenv
  

## Cách sử dụng :

### 1. Sửa đường dẫn tới file ENV tại **Telegram\config\config.go**.
_ Đường dẫn có thể tùy chỉnh theo hệ điều hành đang sử dụng.
_ Thư viện runtime (runtime.GOOS) được sử dụng để tìm ra hệ điều hành đang sử dụng.

### 2. Tạo file ENV :
```bash
# GOOGLE SHEET CONFIG
SPREADSHEETID=''
# Đường dẫn cho file secret xác thực với google sheet
LINUXGGSSCRPATH='' 
WINDOWSGGSSCRPATH=''
# TELEGRAM CONFIG
TELEGRAMBOTAPIKEY=''
# NAMEDOTCOM CONFIG
NAMEDOMAIN=''
NAMEUSRNAME=''
NAMEAPIKEY=''
# REBRANDLY CONFIG
REBRANDLYAPIKEY=''
REBRANDLYDOMAINID=''

```

## Tính năng :

### 1. Tạo forward và short link :

_ Cú pháp : /createfwshortlink [range]

- ví dụ : /createfwshortlink 1001:2002.
- yêu cầu số sau phải lớn hơn số trước.

_ Cần chỉnh sửa tên các cột ;

- name.com handler : cột có chứa store link và temp forward link dùng trong name.com :
```go
func ForwardLinks(inputRange string) (forwardResult []string, successForwardCount int, errorForwardCount int) {
    //
    ...
    //
    //Column assign
    storeLinksCol := "T"
    tempForwardLinksCol := "U"
    //
    ...
    //
}
```

- rebrandly.com handler : cột có chứa slashtag dùng để tạo short links :
```go
func CreateShortLinkRebrandly(inputRange string, inputFwdLinks []string) (shortLinkResult []string, successCount int, errorCount int, usedCount int) {
    //
    ...
    //
    //Column assign
    slashTagCol := "W"
    //
    ...
    //
}
```

_ Nếu đầu vào >= 4 dòng thì chương trình sẽ chạy theo 2 luồng, còn lại là 1 luồng.

_ Chương trình sẽ đợi tạo forward links xong rồi mới tạo short links, forward links được tạo là input để short link.

_ Thời gian chạy chậm do phải đợi khá lâu để lấy data từ google sheet.

### 2. Tính năng có thể phát triển :

- [ ] Tạo forward link.

- [ ] Tạo short links.