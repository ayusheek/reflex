package clients

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func SendHeadlessRequest(url string, showBrowser bool) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var resp string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate("document.documentElement.outerHTML", &resp),
	)

	if err != nil {
		return "", fmt.Errorf("error extracting dom: %s", &err)
	}

	return resp, nil
}
