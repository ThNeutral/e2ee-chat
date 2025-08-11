package utils

import "chat/client/entities"

func InputChangeHandler(this *entities.RectangleWidget, text []rune) {
	for _, char := range text {
		if char == 127 {
			length := len(this.Text)
			if length != 0 {
				this.Text = this.Text[:length-1]
			}
			continue
		}

		this.Text += string(char)
	}
}
