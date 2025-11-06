package processor

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type TextProcessor struct {
	rules map[string]RuleHandler
}

type RuleHandler func([]string, int) []string

func New() *TextProcessor {
	p := &TextProcessor{
		rules: make(map[string]RuleHandler),
	}
	p.registerRules()
	return p
}

func (p *TextProcessor) Process(text string) string {
	words := strings.Fields(text)
	result := p.processRules(words)
	return p.formatOutput(strings.Join(result, " "))
}

func (p *TextProcessor) processRules(words []string) []string {
	var result []string
	
	for i := 0; i < len(words); i++ {
		word := words[i]
		
		if rule := p.extractRule(words, i); rule != "" {
			result = p.applyRule(rule, result)
			if strings.Contains(rule, ",") {
				i++ // Skip next word for range rules
			}
			continue
		}
		
		if p.needsArticleCorrection(word, words, i) {
			result = append(result, p.correctArticle(word))
			continue
		}
		
		result = append(result, word)
	}
	
	return result
}

func (p *TextProcessor) extractRule(words []string, i int) string {
	word := words[i]
	if !strings.HasPrefix(word, "(") {
		return ""
	}
	
	if strings.HasSuffix(word, ")") {
		return word
	}
	
	if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
		return word + " " + words[i+1]
	}
	
	return ""
}

func (p *TextProcessor) applyRule(rule string, result []string) []string {
	rule = strings.Trim(rule, "()")
	rule = strings.ReplaceAll(rule, " ", "")
	parts := strings.Split(rule, ",")
	
	command := strings.TrimSpace(parts[0])
	count := 1
	if len(parts) > 1 {
		if c, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
			count = c
		}
	}
	
	if handler, exists := p.rules[command]; exists {
		return handler(result, count)
	}
	
	return result
}

func (p *TextProcessor) registerRules() {
	p.rules["hex"] = p.convertHex
	p.rules["bin"] = p.convertBin
	p.rules["up"] = p.toUpper
	p.rules["low"] = p.toLower
	p.rules["cap"] = p.capitalize
}

func (p *TextProcessor) convertHex(result []string, _ int) []string {
	if len(result) == 0 {
		return result
	}
	
	for i := len(result) - 1; i >= 0; i-- {
		if val, err := strconv.ParseInt(result[i], 16, 64); err == nil {
			result[i] = strconv.FormatInt(val, 10)
			break
		}
	}
	return result
}

func (p *TextProcessor) convertBin(result []string, _ int) []string {
	if len(result) == 0 {
		return result
	}
	
	for i := len(result) - 1; i >= 0; i-- {
		if val, err := strconv.ParseInt(result[i], 2, 64); err == nil {
			result[i] = strconv.FormatInt(val, 10)
			break
		}
	}
	return result
}

func (p *TextProcessor) toUpper(result []string, count int) []string {
	return p.transformWords(result, count, strings.ToUpper)
}

func (p *TextProcessor) toLower(result []string, count int) []string {
	return p.transformWords(result, count, strings.ToLower)
}

func (p *TextProcessor) capitalize(result []string, count int) []string {
	return p.transformWords(result, count, func(s string) string {
		return strings.Title(strings.ToLower(s))
	})
}

func (p *TextProcessor) transformWords(result []string, count int, transform func(string) string) []string {
	words := 0
	for i := len(result) - 1; i >= 0 && words < count; i-- {
		if result[i] != "" {
			result[i] = transform(result[i])
			words++
		}
	}
	return result
}

func (p *TextProcessor) needsArticleCorrection(word string, words []string, i int) bool {
	return (word == "a" || word == "A") && i+1 < len(words) && p.startsWithVowelOrH(words[i+1])
}

func (p *TextProcessor) correctArticle(word string) string {
	if word == "A" {
		return "An"
	}
	return "an"
}

func (p *TextProcessor) startsWithVowelOrH(word string) bool {
	if len(word) == 0 {
		return false
	}
	first := unicode.ToLower(rune(word[0]))
	return first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' || first == 'h'
}

func (p *TextProcessor) formatOutput(text string) string {
	// Group punctuation
	text = regexp.MustCompile(`\s*\.\s*\.\s*\.\s*`).ReplaceAllString(text, "...")
	text = regexp.MustCompile(`\s*([!?])\s*([!?])\s*`).ReplaceAllString(text, "$1$2")
	
	// Fix punctuation spacing
	text = regexp.MustCompile(`\s+([,.!?:;])`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`([,.!?:;])([^\s,.!?:;])`).ReplaceAllString(text, "$1 $2")
	text = regexp.MustCompile(`(\.{3}|[!?]{2})([^\s])`).ReplaceAllString(text, "$1 $2")
	
	// Fix quotes
	return regexp.MustCompile(`'\s*([^']*?)\s*'`).ReplaceAllStringFunc(text, func(match string) string {
		content := strings.Trim(match, "'")
		content = strings.TrimSpace(content)
		return "'" + content + "'"
	})
}