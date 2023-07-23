package encrypt_test

import (
	"fmt"
	"testing"

	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/encrypt"
	"github.com/stretchr/testify/assert"
)

func TestInterleave(t *testing.T) {
	allArgs := []struct {
		in  string
		out string
	}{
		{"Hello, World!", "H!edlllroo,W "},
		{"We're ¼ of the way there, so ¾ is remaining.", "W.eg'nrien i¼a moefr  tshie  ¾w aoys  t,heer"},
		{"64² = 4096", "6649²0 4= "},
		{"© FÒÖ BÃR BÅZ 2014", "©4 1F0Ò2Ö  ZBÅÃBR "},
		{"Öxxö Xööx \"Lëïth Säë\" - \"Ÿ\"", "Ö\"xŸx\"ö  -X ö\"öëxä S\" Lhëtï"},
		{"Padded with 0xFFÿÿÿÿÿÿÿÿ", "Pÿaÿdÿdÿeÿdÿ ÿwÿiFtFhx 0"},
		{"This string contains NUL\x00 (value 0) and a € (value 128)", "T)h8i2s1  seturlianvg(  c€o nat adinnas  )N0U Le\x00u l(av"},
	}

	for _, arg := range allArgs {
		t.Run(fmt.Sprintf("%s should interleave to %s", arg.in, arg.out), func(t *testing.T) {
			interleaved := encrypt.Interleave(data.BytesFromString(arg.in))
			assert.Equal(t, arg.out, data.StringFromBytes(interleaved))
		})
	}
}

func TestDeinterleave(t *testing.T) {
	allArgs := []struct {
		in  string
		out string
	}{
		{"Hello, World!", "Hlo ol!drW,le"},
		{"We're ¼ of the way there, so ¾ is remaining.", "W'e¼o h a hr,s  srmiig.nnae i¾o eetywetf  re"},
		{"64² = 4096", "6²=4960  4"},
		{"© FÒÖ BÃR BÅZ 2014", "©FÖBRBZ2140 Å Ã Ò "},
		{"Öxxö Xööx \"Lëïth Säë\" - \"Ÿ\"", "Öx öx\"ët ä\"-\"\"Ÿ  ëShïL öXöx"},
		{"Padded with 0xFFÿÿÿÿÿÿÿÿ", "Pde ih0FÿÿÿÿÿÿÿÿFx twdda"},
		{"This string contains NUL\x00 (value 0) and a € (value 128)", "Ti tigcnan U\x00(au )ada€(au 2)81elv   n 0elv LNsito nrssh"},
	}

	for _, arg := range allArgs {
		t.Run(fmt.Sprintf("%s should deinterleave to %s", arg.in, arg.out), func(t *testing.T) {
			deinterleaved := encrypt.Deinterleave(data.BytesFromString(arg.in))
			assert.Equal(t, arg.out, data.StringFromBytes(deinterleaved))
		})
	}
}

func TestFlipMsb(t *testing.T) {
	allArgs := []struct {
		in  string
		out string
	}{
		{"Hello, World!", "Èåììï¬\u00A0×ïòìä¡"},
		{"We're ¼ of the way there, so ¾ is remaining.", "×å§òå\u00A0<\u00A0ïæ\u00A0ôèå\u00A0÷áù\u00A0ôèåòå¬\u00A0óï\u00A0>\u00A0éó\u00A0òåíáéîéîç®"},
		{"64² = 4096", "¶´2\u00A0½\u00A0´°¹¶"},
		{"© FÒÖ BÃR BÅZ 2014", ")\u00A0ÆRV\u00A0ÂCÒ\u00A0ÂEÚ\u00A0²°±´"},
		{"Öxxö Xööx \"Lëïth Säë\" - \"Ÿ\"", "Vøøv\u00A0Øvvø\u00A0¢Ìkoôè\u00A0Ódk¢\u00A0\u00AD\u00A0¢\u001F¢"},
		{"Padded with 0xFFÿÿÿÿÿÿÿÿ", "Ðáääåä\u00A0÷éôè\u00A0°øÆÆ\u007F\u007F\u007F\u007F\u007F\u007F\u007F\u007F"},
		{"This string contains NUL\x00 (value 0) and a € (value 128)", "Ôèéó\u00A0óôòéîç\u00A0ãïîôáéîó\u00A0ÎÕÌ\x00\u00A0¨öáìõå\u00A0°©\u00A0áîä\u00A0á\u00A0€\u00A0¨öáìõå\u00A0±²¸©"},
	}

	for _, arg := range allArgs {
		t.Run(fmt.Sprintf("%s should flip msb to %s", arg.in, arg.out), func(t *testing.T) {
			flipped := encrypt.FlipMsb(data.BytesFromString(arg.in))
			assert.Equal(t, arg.out, data.StringFromBytes(flipped))
		})
	}
}

func TestSwapMultiples(t *testing.T) {
	allArgs := []struct {
		in  string
		out string
	}{
		{"Hello, World!", "Heoll, lroWd!"},
		{"We're ¼ of the way there, so ¾ is remaining.", "Wer'e ¼ fo the way there, so ¾ is remaining."},
		{"64² = 4096", "64² = 4690"},
		{"© FÒÖ BÃR BÅZ 2014", "© FÒÖ ÃBR BÅZ 2014"},
		{"Öxxö Xööx \"Lëïth Säë\" - \"Ÿ\"", "Ööxx Xxöö \"Lëïth Säë\" - \"Ÿ\""},
		{"Padded with 0xFFÿÿÿÿÿÿÿÿ", "Padded with x0FFÿÿÿÿÿÿÿÿ"},
		{"This string contains NUL\x00 (value 0) and a € (value 128)", "This stirng ocntains NUL\x00 (vaule 0) and a € (vaule 128)"},
	}

	for _, arg := range allArgs {
		t.Run(fmt.Sprintf("%s should swap multiples to %s", arg.in, arg.out), func(t *testing.T) {
			swapped, err := encrypt.SwapMultiples(data.BytesFromString(arg.in), 3)
			assert.NoError(t, err)
			assert.Equal(t, arg.out, data.StringFromBytes(swapped))
		})
	}
}
