# CLI Pix
cli-pix is a powerful CLI tool to convert images across formats with support for Windows, Linux, and Mac. It features parallel processing, output folder management, original file deletion, and flexible format selection, enabling streamlined automation with easy CLI flags.

Written in **Go** cause we all love speed 

- ⚡ Bulk image conversion with parallelism

- 📁 Specify output folder

- 🧹 Delete original files post-conversion (optional)

- 🧾 Choose output format (png, webp, jpg, etc.)

- 🛠️ Designed to be modular and extensible


## Installation

###  Linux

 Download the latest release

`wget https://github.com/23Aryansharma11/cli-pix/releases/download/v1.0.0/cli-pix_linux_amd64`

Make it executable

`chmod +x cli-pix_linux_amd64`

Move to a directory in your PATH

`sudo mv cli-pix_linux_amd64 /usr/local/bin/cli-pix`



### For Windows and Mac
Windows and Mac binaries will be available in upcoming releases.

## Build From source

Clone the repo

`git clone https://github.com/23Aryansharma11/cli-pix.git`

`cd cli-pix`


Create a binary named cli-pix in the current folder

`go build -o cli-pix .`

Test 

`./cli-pix`

Move to global and make avaliable anywhere

-  Linux
  `sudo mv cli-pix /usr/local/bin/cli-pix`

- Mac

`sudo mv cli-pix /usr/local/bin/cli-pix`

`chmod +x /usr/local/bin/cli-pix`

- Windows
`go build -o cli-pix.exe .`

`Move-Item -Path cli-pix.exe -Destination "$env:USERPROFILE\AppData\Local\Microsoft\WindowsApps"`
