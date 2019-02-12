package models

import (
	"github.com/AntoineAugusti/moduluschecking/helpers"
)

// Represents a UK bank account
type BankAccount struct {
	SortCode      string
	AccountNumber string
}

// The sort code has an integers slice
func (b BankAccount) SortCodeSlice() []int {
	return helpers.StringToIntSlice(b.SortCode)
}

// The account number has an integers slice
func (b BankAccount) AccountNumberSlice() []int {
	return helpers.StringToIntSlice(b.AccountNumber)
}

// Merge the sort code and the account number of a
// bank account into a single slice
func (b BankAccount) MergeAccountDetails() []int {
	return append(b.SortCodeSlice(), b.AccountNumberSlice()...)
}

// Get the integer value from a letter, according to the defined code:
// Letters between u and z select a digit from the sort code
// Letters between a and h select a digit from the account number
func (b BankAccount) NumberAtPosition(letter string) int {
	nb := b.MergeAccountDetails()
	switch {
	case letter == "u":
		return nb[0]
	case letter == "v":
		return nb[1]
	case letter == "w":
		return nb[2]
	case letter == "x":
		return nb[3]
	case letter == "y":
		return nb[4]
	case letter == "z":
		return nb[5]
	case letter == "a":
		return nb[6]
	case letter == "b":
		return nb[7]
	case letter == "c":
		return nb[8]
	case letter == "d":
		return nb[9]
	case letter == "e":
		return nb[10]
	case letter == "f":
		return nb[11]
	case letter == "g":
		return nb[12]
	case letter == "h":
		return nb[13]
	}

	panic("Unknow letter")
}

// Create a BankAccount structure from a sort code and an account number
func CreateBankAccount(sortCode, accountNumber string) BankAccount {
	accountNumber = helpers.RemoveDashes(accountNumber)

	if len(accountNumber) > 10 || len(accountNumber) < 6 {
		panic("Can handle account number only between 6 and 10 digits")
	}

	if len(sortCode) != 6 {
		panic("Expected a 6 digits sort code")
	}

	switch {
	case (len(accountNumber)) == 10:
		if isCooperativeBankSortCode(sortCode) {
			accountNumber = accountNumber[:8]
		} else {
			accountNumber = accountNumber[len(accountNumber)-8:]
		}
	case len(accountNumber) == 9:
		sortCode = sortCode[:5] + string(accountNumber[0])
		accountNumber = accountNumber[len(accountNumber)-8:]
	case len(accountNumber) <= 7:
		accountNumber = helpers.AddLeadingZeros(accountNumber, 8)
	}

	return BankAccount{
		SortCode:      sortCode,
		AccountNumber: accountNumber,
	}
}

// Check if a sort code is used by the Co-Operative Bank PLC
// The list of sort codes is taken from http://sortcode.a1feeds.com/sortcode/the-co-operative-bank-plc-sort-codes/
func isCooperativeBankSortCode(sortCode string) bool {
	coopSortCodes := map[string]bool{
		"080050": true,
		"080051": true,
		"080052": true,
		"080053": true,
		"080100": true,
		"080211": true,
		"080228": true,
		"080299": true,
		"080308": true,
		"085501": true,
		"085502": true,
		"085503": true,
		"085504": true,
		"085505": true,
		"085506": true,
		"085507": true,
		"085508": true,
		"085509": true,
		"085510": true,
		"085511": true,
		"085512": true,
		"085513": true,
		"085514": true,
		"085515": true,
		"085516": true,
		"085517": true,
		"085518": true,
		"085519": true,
		"085521": true,
		"085523": true,
		"085524": true,
		"085525": true,
		"085526": true,
		"085527": true,
		"085528": true,
		"085529": true,
		"085530": true,
		"085531": true,
		"085536": true,
		"085537": true,
		"085538": true,
		"085539": true,
		"085540": true,
		"085541": true,
		"085543": true,
		"085544": true,
		"085545": true,
		"085546": true,
		"085547": true,
		"085548": true,
		"085549": true,
		"085550": true,
		"085551": true,
		"085556": true,
		"085560": true,
		"085580": true,
		"085593": true,
		"085599": true,
		"086020": true,
		"086090": true,
		"086109": true,
		"086117": true,
		"087062": true,
		"087074": true,
		"087077": true,
		"087078": true,
		"087081": true,
		"087082": true,
		"087085": true,
		"087089": true,
		"087093": true,
		"087094": true,
		"087097": true,
		"087147": true,
		"087162": true,
		"087163": true,
		"087166": true,
		"087167": true,
		"087171": true,
		"087174": true,
		"087175": true,
		"087178": true,
		"087186": true,
		"087191": true,
		"087194": true,
		"087197": true,
		"087198": true,
		"087204": true,
		"087205": true,
		"087207": true,
		"087244": true,
		"087264": true,
		"087267": true,
		"087271": true,
		"087275": true,
		"087278": true,
		"087279": true,
		"087286": true,
		"087287": true,
		"087291": true,
		"087294": true,
		"087295": true,
		"087298": true,
		"087341": true,
		"087361": true,
		"087364": true,
		"087368": true,
		"087372": true,
		"087376": true,
		"087379": true,
		"087384": true,
		"087387": true,
		"087392": true,
		"087395": true,
		"087396": true,
		"087430": true,
		"087449": true,
		"087453": true,
		"087461": true,
		"087465": true,
		"087468": true,
		"087472": true,
		"087473": true,
		"087476": true,
		"087477": true,
		"087481": true,
		"087488": true,
		"087496": true,
		"087499": true,
		"087538": true,
		"087561": true,
		"087562": true,
		"087566": true,
		"087569": true,
		"087577": true,
		"087588": true,
		"087589": true,
		"087596": true,
		"087597": true,
		"087600": true,
		"087623": true,
		"087627": true,
		"087635": true,
		"087669": true,
		"087677": true,
		"087678": true,
		"087686": true,
		"087693": true,
		"087694": true,
		"087701": true,
		"087732": true,
		"087744": true,
		"087752": true,
		"087763": true,
		"087767": true,
		"087771": true,
		"087774": true,
		"087775": true,
		"087779": true,
		"087787": true,
		"087791": true,
		"087795": true,
		"087798": true,
		"087833": true,
		"087841": true,
		"087852": true,
		"087864": true,
		"087868": true,
		"087871": true,
		"087872": true,
		"087879": true,
		"087887": true,
		"087895": true,
		"087898": true,
		"087899": true,
		"087906": true,
		"087922": true,
		"087937": true,
		"087961": true,
		"087968": true,
		"087976": true,
		"087984": true,
		"087987": true,
		"087988": true,
		"088002": true,
		"089000": true,
		"089001": true,
		"089002": true,
		"089003": true,
		"089004": true,
		"089005": true,
		"089006": true,
		"089007": true,
		"089008": true,
		"089009": true,
		"089010": true,
		"089011": true,
		"089012": true,
		"089013": true,
		"089014": true,
		"089015": true,
		"089016": true,
		"089017": true,
		"089018": true,
		"089019": true,
		"089020": true,
		"089021": true,
		"089022": true,
		"089023": true,
		"089024": true,
		"089025": true,
		"089026": true,
		"089027": true,
		"089028": true,
		"089029": true,
		"089030": true,
		"089031": true,
		"089032": true,
		"089033": true,
		"089034": true,
		"089035": true,
		"089036": true,
		"089037": true,
		"089038": true,
		"089039": true,
		"089040": true,
		"089041": true,
		"089042": true,
		"089043": true,
		"089044": true,
		"089045": true,
		"089046": true,
		"089047": true,
		"089048": true,
		"089049": true,
		"089050": true,
		"089051": true,
		"089052": true,
		"089053": true,
		"089054": true,
		"089055": true,
		"089056": true,
		"089057": true,
		"089058": true,
		"089059": true,
		"089060": true,
		"089061": true,
		"089062": true,
		"089063": true,
		"089064": true,
		"089065": true,
		"089066": true,
		"089067": true,
		"089068": true,
		"089069": true,
		"089070": true,
		"089071": true,
		"089072": true,
		"089073": true,
		"089074": true,
		"089075": true,
		"089076": true,
		"089077": true,
		"089078": true,
		"089079": true,
		"089080": true,
		"089081": true,
		"089082": true,
		"089083": true,
		"089084": true,
		"089085": true,
		"089086": true,
		"089087": true,
		"089088": true,
		"089089": true,
		"089090": true,
		"089091": true,
		"089092": true,
		"089093": true,
		"089094": true,
		"089095": true,
		"089096": true,
		"089097": true,
		"089098": true,
		"089099": true,
		"089100": true,
		"089101": true,
		"089102": true,
		"089103": true,
		"089104": true,
		"089105": true,
		"089201": true,
		"089202": true,
		"089203": true,
		"089204": true,
		"089205": true,
		"089206": true,
		"089207": true,
		"089208": true,
		"089209": true,
		"089210": true,
		"089211": true,
		"089212": true,
		"089213": true,
		"089214": true,
		"089215": true,
		"089216": true,
		"089217": true,
		"089218": true,
		"089219": true,
		"089221": true,
		"089222": true,
		"089223": true,
		"089224": true,
		"089225": true,
		"089226": true,
		"089227": true,
		"089228": true,
		"089229": true,
		"089230": true,
		"089231": true,
		"089232": true,
		"089233": true,
		"089234": true,
		"089235": true,
		"089236": true,
		"089237": true,
		"089238": true,
		"089240": true,
		"089241": true,
		"089242": true,
		"089243": true,
		"089244": true,
		"089246": true,
		"089247": true,
		"089248": true,
		"089249": true,
		"089250": true,
		"089252": true,
		"089253": true,
		"089254": true,
		"089255": true,
		"089256": true,
		"089261": true,
		"089262": true,
		"089263": true,
		"089266": true,
		"089267": true,
		"089269": true,
		"089270": true,
		"089272": true,
		"089273": true,
		"089274": true,
		"089275": true,
		"089276": true,
		"089278": true,
		"089279": true,
		"089280": true,
		"089281": true,
		"089282": true,
		"089283": true,
		"089284": true,
		"089285": true,
		"089286": true,
		"089287": true,
		"089288": true,
		"089289": true,
		"089290": true,
		"089291": true,
		"089292": true,
		"089293": true,
		"089294": true,
		"089295": true,
		"089296": true,
		"089297": true,
		"089298": true,
		"089299": true,
		"089300": true,
		"089301": true,
		"089302": true,
		"089303": true,
		"089400": true,
		"089401": true,
		"089402": true,
		"089403": true,
		"089404": true,
		"089405": true,
		"089406": true,
		"089407": true,
		"089408": true,
		"089409": true,
		"089410": true,
		"089411": true,
		"089412": true,
		"839125": true,
		"839126": true,
	}

	_, isPresent := coopSortCodes[sortCode]

	return isPresent
}
