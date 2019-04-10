package slack

// https://api.slack.com/reference/messaging/block-elements

const (
	metImage      MessageElementType = "image"
	metButton     MessageElementType = "button"
	metOverflow   MessageElementType = "overflow"
	metDatepicker MessageElementType = "datepicker"
	metSelect     MessageElementType = "static_select"

	mixedElementImage MixedElementType = "mixed_image"
	mixedElementText  MixedElementType = "mixed_text"
)

type MessageElementType string
type MixedElementType string

// BlockElement defines an interface that all block element types should implement.
type BlockElement interface {
	elementType() MessageElementType
}

type MixedElement interface {
	mixedElementType() MixedElementType
}

type Accessory struct {
	ImageElement      *ImageBlockElement
	ButtonElement     *ButtonBlockElement
	OverflowElement   *OverflowBlockElement
	DatePickerElement *DatePickerBlockElement
	SelectElement     *SelectBlockElement
}

// NewAccessory returns a new Accessory for a given block element
func NewAccessory(element BlockElement) *Accessory {
	switch element.(type) {
	case *ImageBlockElement:
		return &Accessory{ImageElement: element.(*ImageBlockElement)}
	case *ButtonBlockElement:
		return &Accessory{ButtonElement: element.(*ButtonBlockElement)}
	case *OverflowBlockElement:
		return &Accessory{OverflowElement: element.(*OverflowBlockElement)}
	case *DatePickerBlockElement:
		return &Accessory{DatePickerElement: element.(*DatePickerBlockElement)}
	case *SelectBlockElement:
		return &Accessory{SelectElement: element.(*SelectBlockElement)}
	}

	return nil
}

// BlockElements is a convenience struct defined to allow dynamic unmarshalling of
// the "elements" value in Slack's JSON response, which varies depending on BlockElement type
type BlockElements struct {
	BlockElementSet []BlockElement `json:"element"`
}

// ImageBlockElement An element to insert an image - this element can be used
// in section and context blocks only. If you want a block with only an image
// in it, you're looking for the image block.
//
// More Information: https://api.slack.com/reference/messaging/block-elements#image
type ImageBlockElement struct {
	Type     MessageElementType `json:"type"`
	ImageURL string             `json:"image_url"`
	AltText  string             `json:"alt_text"`
}

func (s ImageBlockElement) elementType() MessageElementType {
	return s.Type
}

func (s ImageBlockElement) mixedElementType() MixedElementType {
	return mixedElementImage
}

// NewImageBlockElement returns a new instance of an image block element
func NewImageBlockElement(imageURL, altText string) *ImageBlockElement {
	return &ImageBlockElement{
		Type:     metImage,
		ImageURL: imageURL,
		AltText:  altText,
	}
}

// ButtonBlockElement defines an interactive element that inserts a button. The
// button can be a trigger for anything from opening a simple link to starting
// a complex workflow.
//
// More Information: https://api.slack.com/reference/messaging/block-elements#button
type ButtonBlockElement struct {
	Type     MessageElementType       `json:"type,omitempty"`
	Text     *TextBlockObject         `json:"text"`
	ActionID string                   `json:"action_id,omitempty"`
	URL      string                   `json:"url,omitempty"`
	Value    string                   `json:"value,omitempty"`
	Confirm  *ConfirmationBlockObject `json:"confirm,omitempty"`
}

func (s ButtonBlockElement) elementType() MessageElementType {
	return s.Type
}

// NewButtonBlockElement returns an instance of a new button element to be used within a block
func NewButtonBlockElement(actionID, value string, text *TextBlockObject) *ButtonBlockElement {
	return &ButtonBlockElement{
		Type:     metButton,
		ActionID: actionID,
		Text:     text,
		Value:    value,
	}
}

// SelectBlockElement defines the simplest form of select menu, with a static list
// of options passed in when defining the element.
//
// More Information: https://api.slack.com/reference/messaging/block-elements#select
type SelectBlockElement struct {
	Type          string                    `json:"type,omitempty"`
	Placeholder   *TextBlockObject          `json:"placeholder,omitempty"`
	ActionID      string                    `json:"action_id,omitempty"`
	Options       []*OptionBlockObject      `json:"options,omitempty"`
	OptionGroups  []*OptionGroupBlockObject `json:"option_groups,omitempty"`
	InitialOption *OptionBlockObject        `json:"initial_option,omitempty"`
	Confirm       *ConfirmationBlockObject  `json:"confirm,omitempty"`
}

func (s SelectBlockElement) elementType() MessageElementType {
	return MessageElementType(s.Type)
}

// NewOptionsSelectBlockElement returns a new instance of SelectBlockElement for use with
// the Options object only.
func NewOptionsSelectBlockElement(optType string, placeholder *TextBlockObject, actionID string, options ...*OptionBlockObject) *SelectBlockElement {
	return &SelectBlockElement{
		Type:        optType,
		Placeholder: placeholder,
		ActionID:    actionID,
		Options:     options,
	}
}

// NewOptionsGroupSelectBlockElement returns a new instance of SelectBlockElement for use with
// the Options object only.
func NewOptionsGroupSelectBlockElement(
	optType string,
	placeholder *TextBlockObject,
	actionID string,
	optGroups ...*OptionGroupBlockObject,
) *SelectBlockElement {
	return &SelectBlockElement{
		Type:         optType,
		Placeholder:  placeholder,
		ActionID:     actionID,
		OptionGroups: optGroups,
	}
}

// OverflowBlockElement defines the fields needed to use an overflow element.
// And Overflow Element is like a cross between a button and a select menu -
// when a user clicks on this overflow button, they will be presented with a
// list of options to choose from.
//
// More Information: https://api.slack.com/reference/messaging/block-elements#overflow
type OverflowBlockElement struct {
	Type     MessageElementType       `json:"type"`
	ActionID string                   `json:"action_id,omitempty"`
	Options  []*OptionBlockObject     `json:"options"`
	Confirm  *ConfirmationBlockObject `json:"confirm,omitempty"`
}

func (s OverflowBlockElement) elementType() MessageElementType {
	return s.Type
}

// NewOverflowBlockElement returns an instance of a new Overflow Block Element
func NewOverflowBlockElement(actionID string, options ...*OptionBlockObject) *OverflowBlockElement {
	return &OverflowBlockElement{
		Type:     metOverflow,
		ActionID: actionID,
		Options:  options,
	}
}

// DatePickerBlockElement defines an element which lets users easily select a
// date from a calendar style UI. Date picker elements can be used inside of
// section and actions blocks.
//
// More Information: https://api.slack.com/reference/messaging/block-elements#datepicker
type DatePickerBlockElement struct {
	Type        MessageElementType       `json:"type"`
	ActionID    string                   `json:"action_id"`
	Placeholder *TextBlockObject         `json:"placeholder,omitempty"`
	InitialDate string                   `json:"initial_date,omitempty"`
	Confirm     *ConfirmationBlockObject `json:"confirm,omitempty"`
}

func (s DatePickerBlockElement) elementType() MessageElementType {
	return s.Type
}

// NewDatePickerBlockElement returns an instance of a date picker element
func NewDatePickerBlockElement(actionID string) *DatePickerBlockElement {
	return &DatePickerBlockElement{
		Type:     metDatepicker,
		ActionID: actionID,
	}
}
