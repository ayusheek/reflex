package clients

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func SendHeadlessRequest(url string, showBrowser bool) (string, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", showBrowser))
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var resp string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate("document.documentElement.outerHTML", &resp),
	)

	if err != nil {
		return "", fmt.Errorf("error extracting dom: %s", err)
	}

	return resp, nil
}
