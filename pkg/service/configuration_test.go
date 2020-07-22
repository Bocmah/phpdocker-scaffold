package service_test

import (
	"github.com/Bocmah/phpdocker-scaffold/pkg/service"
	"reflect"
	"testing"
)

func TestLoadConfigFromFile(t *testing.T) {
	got, err := service.LoadConfigFromFile("testdata/test.yaml")

	if err != nil {
		t.Errorf("Got error when loading correct config. Error - %v, Value - %v", err, got)
		return
	}

	want := &service.Configuration{
		AppName: "docker-scaffold",
		PHP: &service.PHP{
			Version:    "7.4",
			Extensions: []string{"mbstring", "zip", "exif", "pcntl", "gd", "pdo_mysql"},
		},
		Nginx: &service.Nginx{
			Port:               80,
			ServerName:         "test-server",
			FastCGIPassPort:    9000,
			FastCGIReadTimeout: 60,
		},
		NodeJS: &service.NodeJS{
			Version: "10",
		},
		Database: &service.Database{
			System:  service.MySQL,
			Version: "5.7",
			Name:    "test-db",
			Port:    3306,
			Credentials: service.Credentials{
				Username:     "bocmah",
				Password:     "test",
				RootPassword: "testRoot",
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Incorrectly loaded configuration. Want %v, got %v", want, got)
	}
}

func TestConfiguration_IsPresent(t *testing.T) {
	conf := &service.Configuration{}

	services := map[string]bool{
		"php":      false,
		"nodejs":   false,
		"nginx":    false,
		"database": false,
	}

	for s, expectedPresent := range services {
		if conf.IsPresent(s) != expectedPresent {
			t.Errorf("Service %s is present in empty configuration", s)
		}
	}

	conf = &service.Configuration{
		PHP: &service.PHP{
			Version:    "7.4",
			Extensions: []string{"mbstring", "zip", "exif", "pcntl", "gd", "pdo_mysql"},
		},
		Nginx: &service.Nginx{
			Port:               80,
			ServerName:         "docker-scaffold",
			FastCGIPassPort:    9000,
			FastCGIReadTimeout: 60,
		},
		NodeJS: &service.NodeJS{
			Version: "10",
		},
		Database: &service.Database{
			System:  service.MySQL,
			Version: "5.7",
			Name:    "docker-scaffold",
			Port:    3306,
			Credentials: service.Credentials{
				Username:     "bocmah",
				Password:     "test",
				RootPassword: "testRoot",
			},
		},
	}

	services = map[string]bool{
		"php":      true,
		"nodejs":   true,
		"nginx":    true,
		"database": true,
	}

	for s, expectedPresent := range services {
		if conf.IsPresent(s) != expectedPresent {
			t.Errorf("Service %s is expected to be present in configuration %v", s, *conf)
		}
	}
}
