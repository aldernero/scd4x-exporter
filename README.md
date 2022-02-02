# scd4x-exporter
A Prometheus exporter for the Adafruit SCD-40 sensor

This repository provides a collector and systemd service definition so that sensor data from the Adafruit SCD-40/1 can be exported to Prometheus (and then visualized with Grafana).

The exported defines 4 gauges:
- `scd4x_co2` : CO2 concentration in ppm
- `scd4x_temperature` : Temperature in degrees Celsius
- `scd4x_humidity` : Relative humidity in percentage

## Prequisites

This assumes you have a sensor installed in the system and that it's in the periodic measurment state. See the [scd4x](https://github.com/aldernero/scd4x) repo for more details.

## Installation

1. Clone this repository
2. cd into the `scd4x-exporter` directory
3. Build the exporter
The build command below is for a Raspberry Pi with a 64-bit kernel. Adjust accordingly for your OS and architecture.
```shell
GOOS=linux GOARCH=arm64 go build -o scd4x_exporter main.go collector.go
```
4. Copy the `scd4x_exporter` binary to `/usr/local/bin/`
```shell
cp scd4x_exporter /usr/local/bin/
```
You may need to `sudo` the command depending on your user and Linux distribution.
5. Copy the service file to the systemd service directory.
Edit the port in the service file to change the port. By default it uses port 9110 to communicate.
```shell
cp scd4x-exporter.service /etc/systemd/system/
```
6. Enable and start the service
```shell
systemctl enable scd4x-exporter
systemctl start scd4x-exporter
```
You should now see the counters in Prometheus.
