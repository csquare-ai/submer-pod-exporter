package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/csquare-ai/submer-pod-exporter/inputs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

var (
	host    net.IP
	port    int
	apiURL  string
	rootCmd = &cobra.Command{
		Use:   "submer-pod-exporter",
		Short: "Prometheus exporter for Submer smart pod.",
		Run: func(cmd *cobra.Command, args []string) {
			recordMetrics()

			log.Printf("Serving at %s:%d\n", host, port)
			http.Handle("/metrics", promhttp.Handler())
			if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil); err != nil {
				panic(err)
			}
		},
	}
	temperature = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "temperature",
		Help:      "The temperature of the smartpod (°C)",
	})
	consumption = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "consumption",
		Help:      "The consumption of the smartpod (kW)",
	})
	dissipation = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "dissipation",
		Help:      "The dissipation of the smartpod (kW)",
	})
	setpoint = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "setpoint",
		Help:      "The setpoint of the smartpod (°C)",
	})
	mpue = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "mpue",
		Help:      "The mPUE of the smartpod",
	})
	pump1rpm = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "pump1rpm",
		Help:      "The pump1rpm of the smartpod (rotations per minute)",
	})
	pump2rpm = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "pump2rpm",
		Help:      "The pump2rpm of the smartpod (rotations per minute)",
	})
)

func recordMetrics() {
	go func() {
		for {
			func() {
				req, err := http.NewRequest("GET", apiURL, nil)
				if err != nil {
					log.Printf("%+v\n", err)
					return
				}
				ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
				defer cancel()

				req = req.WithContext(ctx)
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Printf("%+v\n", err)
					return
				}
				defer resp.Body.Close()

				data := inputs.RealTime{}
				if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
					log.Printf("%+v\n", err)
					return
				}

				log.Printf("Summary: %+v\n", data.Data)
				temperature.Set(data.Data.Temperature)
				consumption.Set(data.Data.Consumption)
				dissipation.Set(data.Data.Dissipation)
				setpoint.Set(data.Data.Setpoint)
				mpue.Set(data.Data.Mpue)
				pump1rpm.Set(data.Data.Pump1RPM)
				pump2rpm.Set(data.Data.Pump2RPM)
			}()
			time.Sleep(time.Second)
		}
	}()
}

func init() {
	rootCmd.PersistentFlags().IPVar(&host, "host", net.IPv4zero, "listening host")
	rootCmd.PersistentFlags().IntVar(&port, "port", 3000, "listening port")
	rootCmd.PersistentFlags().StringVar(&apiURL, "api-url", "http://localhost/api/realTime", "Submer ssmartpod API URL")
}

func Execute() error {
	return rootCmd.Execute()
}
