# gowall

This is a simple and efficient wallpaper downloader written in Go. It uses the Wallhaven API to fetch and download wallpapers.

## Features

- Download high-quality wallpapers from Wallhaven.
- Filter wallpapers by category, purity, sorting, order, etc.
- Download multiple wallpapers concurrently.

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/M00t3/gowall.git
```

Navigate to the project directory:

```bash
cd wallpaper-downloader
```

Build the project:

```bash
go build
```

for install the project:

```bash
go install
```

## Usage

Run the application:

```bash
./gowall
```

## Configuration

You can configure the application by modifying the `config.go` file. Here are the available options:

- `api_key`: Your Wallhaven API key. (default must be "")
- `category`: The category of wallpapers to download.
- `purity`: The purity of wallpapers to download.
- `sorting`: The sorting method for the wallpapers.
- `order`: The order of the wallpapers.
- `atleast`: The minimum resolution of the wallpapers.
- `topRange`: The top range of the wallpapers.
- `sorting`: The sorting method for the wallpapers.
- `save_directory` : The directory to save the wallpapers.

you can see more details about the options in the [Wallhaven API documentation](https://wallhaven.cc/help/api).

## TODO

- [ ] select category, purity, sorting, etc. from the command line.
- [ ] add a progress bar.
- [ ] optional notification when the download is finished.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the GPL3 License - see the [LICENSE.txt](LICENSE.txt) file for details.
