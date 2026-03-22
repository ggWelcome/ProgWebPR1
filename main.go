package main // головний пакет програми, точка входу

import (
	"fmt"      // пакет для форматованого виводу (використовується для HTML-відповіді)
	"log"      // пакет для логування повідомлень у консоль
	"net/http" // пакет для створення веб-сервера та роботи з HTTP-запитами
	"time"     // пакет для роботи з датою та часом
)

// statusHandler — функція-обробник HTTP-запитів
func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Отримуємо параметр ?mode=... (режим роботи системи)
	mode := r.URL.Query().Get("mode")
	if mode == "" {
		mode = "normal" // якщо параметр не задано — режим "normal"
	}

	// Отримуємо параметр ?interface=... (варіант інтерфейсу)
	ui := r.URL.Query().Get("interface")
	if ui == "" {
		ui = "default" // якщо параметр не задано — інтерфейс "default"
	}

	// Логування запиту у консоль:
	// показуємо IP користувача, шлях сторінки, метод запиту та передані параметри
	log.Printf(
		"[%s] Користувач %s відкрив %s (метод: %s) | Параметри: mode=%s, interface=%s",
		time.Now().Format("02-01-2006 15:04:05"), // час запиту
		r.RemoteAddr,                             // IP-адреса користувача
		r.URL.Path,                               // шлях сторінки (наприклад "/")
		r.Method,                                 // метод запиту (GET/POST)
		mode,                                     // параметр mode
		ui,                                       // параметр interface
	)

	// Формуємо HTML-відповідь для браузера
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // заголовок відповіді
	fmt.Fprintf(w, "<html><head><title>Система розподілу електроенергії (Smart Grid)</title></head><body>")
	fmt.Fprintf(w, "<h1>Стан системи розподілу електроенергії</h1>")
	fmt.Fprintf(w, "<p>Час перевірки: %s</p>", time.Now().Format("15:04:05")) // поточний час

	// Відображення режиму роботи системи залежно від параметра mode
	switch mode {
	case "emergency":
		fmt.Fprintf(w, "<p style='color:red;'>Аварія! Перехід на резервні джерела.</p>")
	case "backup":
		fmt.Fprintf(w, "<p style='color:orange;'>Система працює від резервних джерел.</p>")
	default:
		fmt.Fprintf(w, "<p style='color:green;'>Система працює у штатному режимі.</p>")
	}

	// Варіативність інтерфейсів залежно від параметра interface
	fmt.Fprintf(w, "<hr>") // горизонтальна лінія для відділення блоків
	switch ui {
	case "grid":
		fmt.Fprintf(w, "<h2>Інтерфейс: Smart Grid</h2>")
		fmt.Fprintf(w, "<p>Відображення загального навантаження та генерації відновлюваних джерел.</p>")
	case "substation":
		fmt.Fprintf(w, "<h2>Інтерфейс: Підстанція</h2>")
		fmt.Fprintf(w, "<p>Моніторинг трансформаторів, напруги та струму.</p>")
	case "consumer":
		fmt.Fprintf(w, "<h2>Інтерфейс: Споживачі</h2>")
		fmt.Fprintf(w, "<p>Статистика споживання та пікових навантажень.</p>")
	default:
		fmt.Fprintf(w, "<h2>Інтерфейс: Smart Grid (за замовчуванням)</h2>")
	}
	fmt.Fprintf(w, "</body></html>") // завершення HTML-документа
}

// main — точка входу програми
func main() {
	// Реєструємо обробник для кореневого шляху "/"
	http.HandleFunc("/", statusHandler)

	// Повідомлення у консоль про запуск сервера
	log.Println("Сервер запущено на http://localhost:8080")

	// Запускаємо сервер на порту 8080
	// log.Fatal завершує програму, якщо сервер не зміг стартувати
	log.Fatal(http.ListenAndServe(":8080", nil))
}
