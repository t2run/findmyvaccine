package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"t2findmyvaccinebot/pkg/common"
	"t2findmyvaccinebot/pkg/cowinutils"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// start introduces the bot
func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello"+emoji(`"\xE2\x9C\x8B"`)+", I'm @%s "+emoji(`"\xF0\x9F\x91\xBE"`)+". I <b>will help you to find </b>available vaccine slots."+emoji(`"\xF0\x9F\x92\x89"`), b.User.Username), &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Press me", CallbackData: "start_callback"},
			}},
		},
	})
	if err != nil {
		fmt.Println("failed to send: " + err.Error())
	}
	return nil
}

// startCB edits the start message
func StartCB(b *gotgbot.Bot, ctx *ext.Context) error {
	slotHelp := "\n  " + emoji(`"\x31\xE2\x83\xA3"`) + "Find Slots by PinCode \nPin <pincode>,<DD-MM-YYYY> \n eg: Pin 110012,06-05-2021 \n \n" + emoji(`"\x32\xE2\x83\xA3"`) + "Find Slots by District Code \nFind <districtcode>,<DD-MM-YYYY> \n  eg: Find 140,06-05-2021 \n\n " + emoji(`"\x33\xE2\x83\xA3"`) + "To get your district code \n Dist <statecode> \n  eg: Dist 17 \n\n " + emoji(`"\x34\xE2\x83\xA3"`) + " To get your State Code \n /states \n"
	cb := ctx.Update.CallbackQuery
	cb.Answer(b, nil)
	cb.Message.EditText(b, slotHelp, nil)
	return nil
}

// states replies to a messages with all states and codes
func States(b *gotgbot.Bot, ctx *ext.Context) error {

	var states common.StateList
	//getting from the cache
	errUnMarshal := json.Unmarshal([]byte(InStates), &states)
	if errUnMarshal != nil {
		fmt.Println("failed to marshal: " + errUnMarshal.Error())
		return errUnMarshal
	}
	var sb strings.Builder
	for _, state := range states.States {

		sb.WriteString(state.StateName + " : " + strconv.Itoa(state.StateID) + "\n")
	}
	ctx.EffectiveMessage.Reply(b, sb.String(), nil)
	return nil
}

// KlDist replies to a messages with KL dist and codes
func KlDist(b *gotgbot.Bot, ctx *ext.Context) error {

	var districtList common.DistrictList

	errUnMarshal := json.Unmarshal([]byte(KlDistricts), &districtList)
	if errUnMarshal != nil {
		fmt.Println("failed to marshal: " + errUnMarshal.Error())
		return errUnMarshal
	}
	var sb strings.Builder
	for _, dist := range districtList.Districts {
		sb.WriteString(dist.DistrictName + " : " + strconv.Itoa(dist.DistrictID) + "\n")
	}
	ctx.EffectiveMessage.Reply(b, sb.String(), nil)
	return nil
}

//Handle all requests
func HandleAll(b *gotgbot.Bot, ctx *ext.Context) error {

	slotHelp := "\n  " + emoji(`"\x31\xE2\x83\xA3"`) + "Find Slots by PinCode \nPin <pincode>,<DD-MM-YYYY> \n eg: Pin 110012,06-05-2021 \n \n" + emoji(`"\x32\xE2\x83\xA3"`) + "Find Slots by District Code \nFind <districtcode>,<DD-MM-YYYY> \n  eg: Find 140,06-05-2021 \n\n " + emoji(`"\x33\xE2\x83\xA3"`) + "To get your district code \n Dist <statecode> \n  eg: Dist 17 \n\n " + emoji(`"\x34\xE2\x83\xA3"`) + " To get your State Code \n /states \n"
	message := strings.ToLower(ctx.EffectiveMessage.Text)

	if strings.HasPrefix(message, "dist") {
		return GetDist(b, ctx)
	} else if strings.HasPrefix(message, "find") {
		return AppoinmentByDist(b, ctx)
	} else if strings.HasPrefix(message, "pin") {
		return AppoinmentByPin(b, ctx)
	}
	ctx.EffectiveMessage.Reply(b, slotHelp, nil)
	return nil
}

// AppoinmentByDist replies to a messages with all available slots for the dist code and date
func AppoinmentByDist(b *gotgbot.Bot, ctx *ext.Context) error {

	message := strings.ToLower(ctx.EffectiveMessage.Text)

	message = strings.TrimSpace(strings.Replace(message, "find", "", -1))

	mArray := strings.Split(message, ",")
	var distID, date string
	if len(mArray) == 2 {
		distID = strings.TrimSpace(mArray[0])
		date = strings.TrimSpace(mArray[1])
	} else {
		ctx.EffectiveMessage.Reply(b, emoji(`"\xF0\x9F\x98\xB8"`)+"Invalid input!!! Try this format > Find 140,06-05-2021", nil)
		return nil
	}
	sessions, errGetSession := cowinutils.GetSessionByDist(distID, date)
	if errGetSession != nil {
		fmt.Println("Error Get sessions by district: " + errGetSession.Error())
		return errGetSession
	}

	var sb strings.Builder //strconv.Itoa(dist.DistrictID)
	if len(sessions.Sessions) > 0 {
		for count, slot := range sessions.Sessions {
			sb.WriteString(
				"Available Center # " + strconv.Itoa(count+1) + "\n" +
					" Name: " + slot.Name + "\n" +
					" Address: " + slot.Address + "\n" +
					" Fee Type: " + slot.FeeType + "\n" +
					" Min Age Limit: " + strconv.Itoa(slot.MinAgeLimit) + "\n" +
					" Vaccine Name: " + slot.Vaccine + "\n" +
					" Time Slots: " + fmt.Sprint(slot.Slots) + "\n" +
					" Available Capacity: " + strconv.Itoa(slot.AvailableCapacity) + "\n" +
					"******************************\n")
		}
		sb.WriteString(emoji(`"\xF0\x9F\x92\x89"`) + "Book Now : https://www.cowin.gov.in/home")
	} else {
		sb.WriteString(emoji(`"\xF0\x9F\x98\x9E"`) + "No Slots available")
	}

	ctx.EffectiveMessage.Reply(b, sb.String(), nil)
	return nil
}

// AppoinmentByPin replies to a messages with all available slots for the pin code and date
func AppoinmentByPin(b *gotgbot.Bot, ctx *ext.Context) error {

	message := strings.ToLower(ctx.EffectiveMessage.Text)

	message = strings.TrimSpace(strings.Replace(message, "pin", "", -1))

	mArray := strings.Split(message, ",")
	var pinId, date string
	if len(mArray) == 2 {
		pinId = strings.TrimSpace(mArray[0])
		date = strings.TrimSpace(mArray[1])
	} else {
		ctx.EffectiveMessage.Reply(b, emoji(`"\xF0\x9F\x98\xB8"`)+"Invalid input!!! Try this format > Pin 671310,06-05-2021", nil)
		return nil
	}

	sessions, errGetSession := cowinutils.GetSessionByPin(pinId, date)
	if errGetSession != nil {
		fmt.Println("Error Get sessions by Pin: " + errGetSession.Error())
		return errGetSession
	}

	var sb strings.Builder //strconv.Itoa(dist.DistrictID)
	if len(sessions.Sessions) > 0 {
		for count, slot := range sessions.Sessions {
			sb.WriteString(
				"Available Center # " + strconv.Itoa(count+1) + "\n" +
					" Name: " + slot.Name + "\n" +
					" Address: " + slot.Address + "\n" +
					" Fee Type: " + slot.FeeType + "\n" +
					" MinAgeLimit: " + strconv.Itoa(slot.MinAgeLimit) + "\n" +
					" Vaccine Name: " + slot.Vaccine + "\n" +
					" Time Slots: " + fmt.Sprint(slot.Slots) + "\n" +
					" Available Capacity: " + strconv.Itoa(slot.AvailableCapacity) + "\n" +
					"******************************\n")
		}
		sb.WriteString(emoji(`"\xF0\x9F\x92\x89"`) + "Book Now : https://www.cowin.gov.in/home")
	} else {
		sb.WriteString(emoji(`"\xF0\x9F\x98\x9E"`) + "No Slots available")
	}

	ctx.EffectiveMessage.Reply(b, sb.String(), nil)
	return nil
}

// GetDist replies to a messages with all states and codes
func GetDist(b *gotgbot.Bot, ctx *ext.Context) error {

	message := strings.ToLower(ctx.EffectiveMessage.Text)

	stateID := strings.TrimSpace(strings.Replace(message, "dist", "", -1))

	if len(stateID) == 0 {
		ctx.EffectiveMessage.Reply(b, emoji(`"\xF0\x9F\x98\xB8"`)+"Invalid input!!! Try this format > Dist 17", nil)
	}

	// distList, errGetSession := cowinutils.GetDistricts(stateID)
	// if errGetSession != nil {
	// 	fmt.Println("Error Get sessions by district: " + errGetSession.Error())
	// 	return errGetSession
	// }

	// getting from the cache
	var distString string
	switch stateID {
	case "1":
		distString = state1
	case "2":
		distString = state2
	case "3":
		distString = state3
	case "4":
		distString = state4
	case "5":
		distString = state5
	case "6":
		distString = state6
	case "7":
		distString = state7
	case "8":
		distString = state8
	case "9":
		distString = state9
	case "10":
		distString = state10
	case "11":
		distString = state11
	case "12":
		distString = state12
	case "13":
		distString = state13
	case "14":
		distString = state14
	case "15":
		distString = state15
	case "16":
		distString = state16
	case "17":
		distString = state17
	case "18":
		distString = state18
	case "19":
		distString = state19
	case "20":
		distString = state20
	case "21":
		distString = state21
	case "22":
		distString = state22
	case "23":
		distString = state23
	case "24":
		distString = state24
	case "25":
		distString = state25
	case "26":
		distString = state26
	case "27":
		distString = state27
	case "28":
		distString = state28
	case "29":
		distString = state29
	case "30":
		distString = state30
	case "31":
		distString = state31
	case "32":
		distString = state32
	case "33":
		distString = state33
	case "34":
		distString = state34
	case "35":
		distString = state35
	case "36":
		distString = state36
	case "37":
		distString = state37
	default:
		return States(b, ctx)
	}

	var distList common.DistrictList
	json.Unmarshal([]byte(distString), &distList)

	var sb strings.Builder
	for _, dist := range distList.Districts {
		sb.WriteString(dist.DistrictName + " : " + strconv.Itoa(dist.DistrictID) + "\n")
	}
	ctx.EffectiveMessage.Reply(b, sb.String(), nil)
	return nil
}

func emoji(code string) string {
	value, _ := strconv.Unquote(code)
	return value
}
