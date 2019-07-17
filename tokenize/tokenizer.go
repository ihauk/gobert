// package Tokenization supplies tokenizzation operations for BERT.
// Ports the tokenizer.py capbilites from the core BERT repo
// NOTE: All defintions are related to BERT and may vary from unicode defintions,
// for example, BERT considers '$' punctuation, but unicode does not.
package tokenize

import "github.com/buckhx/gobert/vocab"

// Tokenizer is an interface for chunking a string into it's tokens as per the BERT implematation
type Tokenizer interface {
	Tokenize(text string) (tokens []string)
}

// NewTokenizer returns a new FullTokenizer
func NewTokenizer(voc vocab.Vocab, opts ...Option) Tokenizer {
	tkz := Full{
		Basic:     NewBasic(),
		Wordpiece: NewWordpiece(voc),
	}
	for _, opt := range opts {
		tkz = opt(tkz)
	}
	return tkz
}

// Option alter the behavior of the tokenizer
// TODO add tests for these behavior changes
type Option func(tkz Full) Full

// WithLower will lowercase all input
func WithLower() Option {
	return func(tkz Full) Full {
		tkz.Basic.Lower = true
		return tkz
	}
}

// WithUnkownToken will alter the unkown token from default [UNK]
func WithUnknownToken(unk string) Option {
	return func(tkz Full) Full {
		tkz.Wordpiece.unknownToken = unk
		return tkz
	}
}

// WithMaxChars sets the maximum len of a token to be tokenized, if longer will be labeled as unknown
func WithMaxChars(wc int) Option {
	return func(tkz Full) Full {
		tkz.Wordpiece.maxWordChars = wc
		return tkz
	}
}