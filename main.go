package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

// Tab contém o contexto de uma aba no google chrome
type Tab struct {
	context context.Context
	cancel  context.CancelFunc
}

func main() {

	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("restore-on-startup", false),
		chromedp.Flag("disable-hang-monitor", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-notification", true),
		chromedp.Flag("disable-checker-imaging", true),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("disable-accelerated-2d-canvas", true),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.Flag("safebrowsing-disable-auto-update", true),
		chromedp.Flag("disable-client-side-phishing-detection", true),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), options...)
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	err := chromedp.Run(ctx)
	checkErr(err)

	sent := 0
	for {
		if countTabs(ctx) < 10 {
			tab := newTab(ctx)
			go fillForm(tab)
			sent++
			log.Print(sent)
		}
	}
}

// checkErr testa se existe um error
func checkErr(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// newTab abre uma nova aba no google chrome
func newTab(ctx context.Context) Tab {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	ctx, _ = chromedp.NewContext(ctx)
	return Tab{context: ctx, cancel: cancel}
}

// fillForm preenche e envia um formulário
func fillForm(tab Tab) {
	err := chromedp.Run(tab.context, chromedp.Tasks{
		chromedp.Navigate(`https://grupoboticario.secure.force.com/formularioWeb?customId=CRCBOT`),
		chromedp.WaitVisible(`//*[@id="j_id0:frm"]`),
		chromedp.SendKeys(`#a2Q1U000000SPHjUAO`, "Não sou a Rafaela, parem de enviar cobranças", chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHgUAO`, "Não sou a Rafaela, parem de enviar cobranças", chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHeUAO`, "00000000000", chromedp.ByID),
		chromedp.SetAttributeValue(`#a2Q1U000000SPHeUAO`, "data-role", ""),
		chromedp.SendKeys(`#a2Q1U000000SPHkUAO`, "00000000000", chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHiUAO`, "00000000000", chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHnUAO`, "nao.sou.a.rafaela@nao.sou.a.rafaela.com", chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHlUAO`, "Não sou a Rafaela, parem de enviar cobranças", chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHmUAO`, "Pagamento", chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHfUAO`, "Cobrança Indevida/Cupom Fiscal", chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHhUAO`, "Não sou a Rafaela, parem de enviar cobranças para (00) 0 00000-0000", chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000001dnlJUAQ`, "Não", chromedp.ByID),
		chromedp.Click(`//*[@id="j_id0:frm:j_id11"]/div[2]/div/div`, chromedp.BySearch),
		chromedp.WaitVisible(`body > div > p > i > b`),
	})
	checkErr(err)
	tab.cancel()
}

// countTabs conta quantas abas foram abertas no google chrome
func countTabs(ctx context.Context) int {
	targets, err := chromedp.Targets(ctx)
	checkErr(err)
	result := 0
	for _, t := range targets {
		if t.Type == "page" {
			result++
		}
	}
	return result
}
