package main

import (
	"log"
	"context"
	"github.com/chromedp/chromedp"
)

func main() {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("headless", true),
    chromedp.Flag("disable-gpu", false),
    chromedp.Flag("enable-automation", false),
    chromedp.Flag("disable-extensions", false),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	count := 0
	for {
		err := chromedp.Run(ctx, sendkeys(
			"Não sou a Rafaela, parem de enviar cobranças", 
			"Não sou a Rafaela, parem de enviar cobranças", 
			"00000000000",
			"00000000000",
			"00000000000",
			"nao.sou.a.rafaela@nao.sou.a.rafaela.com",
			"Não sou a Rafaela, parem de enviar cobranças",
			"Pagamento",
			"Cobrança Indevida/Cupom Fiscal",
			"Não sou a Rafaela, parem de enviar cobranças para (31) 9 99742-7487",
			"Não"))
		if err != nil {
			log.Fatal(err)
		}

		count++
		log.Printf("Count: `%d`", count)
	}
}

func sendkeys(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 string) chromedp.Tasks { 
	return chromedp.Tasks{
		chromedp.Navigate(`https://grupoboticario.secure.force.com/formularioWeb?customId=CRCBOT`),
		chromedp.WaitVisible(`//*[@id="j_id0:frm"]`),
		chromedp.SendKeys(`#a2Q1U000000SPHjUAO`, v1, chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHgUAO`, v2, chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHeUAO`, v3, chromedp.ByID),
		chromedp.SetAttributeValue(`#a2Q1U000000SPHeUAO`, "data-role", ""),
		chromedp.SendKeys(`#a2Q1U000000SPHkUAO`, v4, chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHiUAO`, v5, chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHnUAO`, v6, chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHlUAO`, v7, chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHmUAO`, v8, chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHfUAO`, v9, chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000000SPHhUAO`, v10, chromedp.ByID),
		chromedp.SendKeys(`#a2Q1U000001dnlJUAQ`, v11, chromedp.ByID),
		chromedp.Click(`//*[@id="j_id0:frm:j_id11"]/div[2]/div/div`, chromedp.BySearch),
		chromedp.WaitVisible(`body > div > p > i > b`),
	}
}