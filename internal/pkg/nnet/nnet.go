package nnet

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func HostIP() (string, error) {
	var hostIP string

	// Получаем список всех сетевых интерфейсов
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Ошибка при получении сетевых интерфейсов:", err)
		return "", err
	}

	// Перебираем все интерфейсы
	for _, iface := range interfaces {
		// Получаем список всех адресов интерфейса
		addrs, err := iface.Addrs()
		if err != nil {
			log.Println("Ошибка при получении адресов интерфейса:", err)
			continue
		}

		// Игнорируем виртуальные и loopback интерфейсы
		if strings.HasPrefix(iface.Name, "docker") ||
			strings.HasPrefix(iface.Name, "br-") ||
			strings.HasPrefix(iface.Name, "veth") ||
			strings.HasPrefix(iface.Name, "lo") {
			continue
		}

		// Перебираем все адреса интерфейса
		for _, addr := range addrs {
			// Преобразуем адрес в строку
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			// Проверяем, что это не loopback (локальный интерфейс)
			if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				hostIP = ipNet.IP.String()
				break
			}
		}
	}

	return hostIP, nil
}
