package main

import (
	"log"

	"github.com/aldernero/scd4x"
	"github.com/prometheus/client_golang/prometheus"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

type scd4xCollector struct {
	co2         *prometheus.Desc
	temperature *prometheus.Desc
	humidity    *prometheus.Desc
}

func newScd4xCollector() *scd4xCollector {
	return &scd4xCollector{
		co2:         prometheus.NewDesc("scd4x_co2", "CO2 in ppm", nil, nil),
		temperature: prometheus.NewDesc("scd4x_temperature", "Temperature in degrees Celsius", nil, nil),
		humidity:    prometheus.NewDesc("scd4x_humidity", "Relative humidity as a percentage", nil, nil),
	}
}

func (collector *scd4xCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.co2
	ch <- collector.temperature
	ch <- collector.humidity
}

func (collector *scd4xCollector) Collect(ch chan<- prometheus.Metric) {
	_, err := host.Init()
	if err != nil {
		log.Fatalf("Failed to initialize periph: %v", err)
	}
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatalf("Failed while opening bus: %v", err)
	}
	defer bus.Close()
	sensor, err := scd4x.SensorInit(bus, false)
	if err != nil {
		log.Fatal(err)
	}
	data, err := sensor.ReadMeasurement()
	if err != nil {
		log.Fatal(err)
	}
	ch <- prometheus.MustNewConstMetric(collector.co2, prometheus.GaugeValue, float64(data.CO2))
	ch <- prometheus.MustNewConstMetric(collector.temperature, prometheus.GaugeValue, float64(data.Temp))
	ch <- prometheus.MustNewConstMetric(collector.humidity, prometheus.GaugeValue, float64(data.Rh))
}
