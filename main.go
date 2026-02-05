package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	// Setup per la pulizia finale all'uscita (Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		// Sequenza ANSI: Pulisce tutto e resetta il cursore
		fmt.Print("\033[2J\033[H")
		os.Exit(0)
	}()

	jsonData := `{
		"1_popolo_e_famiglia": ["le mamme", "i papà", "i nonni", "i bambini col sorriso", "i figli", "i nipoti", "i ragazzi", "i giovani", "le donne", "gli uomini", "le famiglie", "i padri separati", "le mamme lavoratrici", "la nonna", "la moglie invalida", "i bimbi disabili", "gli anziani soli", "le persone perbene", "gli italiani", "la gente normale", "i cittadini che resistono", "i nostri militanti", "i volontari", "le ragazze", "i genitori", "le comunità", "i vicini di casa", "il popolo delle piazze", "gli italiani dimenticati", "la gente di cuore", "le donne coraggiose", "chi non molla"],
		"2_professioni_e_lavoratori": ["gli artigiani", "i commercianti", "gli agricoltori", "gli allevatori", "i pescatori", "gli imprenditori", "le partite iva", "i liberi professionisti", "gli operai", "i metalmeccanici", "i muratori", "gli idraulici", "gli elettricisti", "i barbieri", "le parrucchiere", "i panettieri", "i ristoratori", "i baristi", "i camionisti", "i tassisti", "i capotreni", "i postini", "i precari", "i disoccupati", "i medici", "gli infermieri", "gli insegnanti", "le maestre", "i ricercatori", "gli studenti", "gli artisti", "le guardie giurate"],
		"3_forze_dell_ordine_e_sicurezza": ["i poliziotti", "i carabinieri", "i finanzieri", "la penitenziaria", "i vigili del fuoco", "le divise", "l'esercito", "i militari", "i pompieri", "la protezione civile", "i vigili urbani", "gli agenti", "chi difende i confini", "chi rischia la vita", "le caserme", "le stazioni di polizia", "la gazzella", "chi lotta contro la mafia", "chi combatte lo spaccio", "chi ferma i delinquenti", "chi protegge le case", "chi ama la divisa", "un eroe quotidiano"],
		"6_tasse_e_fisco": ["la tassa sulla plastica", "la tassa sullo zucchero", "la tassa sul diesel", "la tassa sulla casa", "i balzelli", "le accise", "l'imu", "la tari", "l'irpef", "l'iva", "le cartelle esattoriali", "il fisco nemico", "la burocrazia", "il fisco vampiro", "le micro tasse"],
		"9_cibo_e_spesa": ["i biscotti", "le merendine", "i gelati", "l'insalata in busta", "la frutta fresca", "il pane caldo", "il latte", "le uova", "la carne", "il prosciutto", "il formaggio", "i tortellini", "la pasta", "il vino buono", "il chinotto", "la cedrata", "il caffè", "la passata di pomodoro", "i prodotti locali", "la cucina italiana"],
		"20_slogan_e_frasi_fatte": ["prima gli italiani", "i porti chiusi", "la legittima difesa", "la ruspa", "quota cento", "la flat tax", "la pace fiscale", "la tolleranza zero", "le piazze pulite", "le scuole sicure", "i baci ai rosiconi", "un abbraccio ai gufi", "avanti tutta", "viva la lega", "viva l'italia"]
	}`

	var categorie map[string][]string
	json.Unmarshal([]byte(jsonData), &categorie)

	const art = `######################################################################################
######################################################################################
######################################################################################
#######################+-------########################+....-++++--...-###############
#####################--........---##################+..+###############-.-############
####################-..++++-----++.+##############-.+#####################..##########
###################-..+++++++++++++.+############.-########################+.+########
###################..-+#++++++++++++-###########.+###########################.+#######
###################..-+#++++++++++##-##########.-#############################.+######
###################--+#+++--+++---+#+########## ##############################.-######
###################+-#++++--++---++###########+ ##############################.-######
###################+++#+++++#++++++###########+ ##############################.-######
####################-++++-++#++--+++###########.+#############################.+######
####################++++..----.-.-++############.+###########################.-#######
#####################----+++#+-+---##############.-#########################.-########
###################+#+-.-+++++++--################-.######################- ##########
###################.##+-.++++++-.###################- +################+..############
#################-..+###+.......+-###################.####++######+-...+##############
##############+......+#####-----#+...-+#############.-##-.-##++++#####################
##########-...........+##++##+-#+-..........-######.  ..##############################
#####+.................-#####+##-..-.........#########################################
###-....................+#######-.............-#######################################
###......................#######+..............+######################################
###-......................######+...............######################################
###-......................+#####+...............-#####################################
###+.......................######................+####################################
###-.......... ............-#####.................+###################################
###+........................#####..................+##################################
####........................+####...................+#################################
####.........................####-..........+++.....-#################################
####-........................####-........-+++++.....-################################
#####........................+###-.......+++++-.......################################
#####-.......................-###+-....................###############################
######.............++++#####++++####...................+##############################
######+.............-+++++###++++###+..................+##############################
#######.................++++-----+##+.....--+-.........+##############################
#######+..........................----................-###############################
########-..........................  ...........--+###################################
########-.............................--.--+##########################################
########+.............................+....-##########################################
########-..................................-##########################################
########-............  ........-............##########################################
########...............  ... ...-...........+#########################################
#######-.....................++#+##-........-#########################################
#######......................+##+##+-........#########################################
######-............ .........-+++##+-........#########################################
######+............ ........--+++#++-........+########################################`

	const leftWall, rightWall, centerY = 48, 79, 10
	centerX := leftWall + (rightWall-leftWall)/2

	keys := make([]string, 0, len(categorie))
	for k := range categorie {
		keys = append(keys, k)
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Print("\033[2J")

	for {
		catKey := keys[rand.Intn(len(keys))]
		paroleDellaCat := categorie[catKey]
		numDaMostrare := rand.Intn(2) + 3

		rand.Shuffle(len(paroleDellaCat), func(i, j int) {
			paroleDellaCat[i], paroleDellaCat[j] = paroleDellaCat[j], paroleDellaCat[i]
		})

		if len(paroleDellaCat) < numDaMostrare {
			numDaMostrare = len(paroleDellaCat)
		}
		selezione := paroleDellaCat[:numDaMostrare]

		for _, msg := range selezione {
			fmt.Print("\033[H")

			displayMsg := msg
			if len(msg) > 14 {
				parts := strings.Split(msg, " ")
				if len(parts) > 2 {
					displayMsg = parts[0] + " " + parts[1] + "\n" + strings.Join(parts[2:], " ")
				} else if len(parts) == 2 {
					displayMsg = parts[0] + "\n" + parts[1]
				}
			}

			lines := strings.Split(art, "\n")
			matrix := make([][]rune, len(lines))
			for i, l := range lines {
				matrix[i] = []rune(l)
			}

			msgLines := strings.Split(displayMsg, "\n")
			h := len(msgLines)
			startY := centerY - (h / 2)

			maxW := 0
			for _, l := range msgLines {
				l = strings.TrimSpace(l)
				if len(l) > maxW {
					maxW = len(l)
				}
			}

			whiteWidth := maxW + 4
			areaStartY, areaEndY := startY-1, startY+h
			whiteStartCol := centerX - (whiteWidth / 2)
			whiteEndCol := whiteStartCol + whiteWidth

			for r := areaStartY; r <= areaEndY; r++ {
				for c := whiteStartCol; c < whiteEndCol; c++ {
					if r < 0 || r >= len(matrix) || c < leftWall || c >= rightWall {
						continue
					}
					isTopCorner := r == areaStartY && (c == whiteStartCol || c == whiteEndCol-1)
					isBottomCorner := r == areaEndY && (c == whiteStartCol || c == whiteEndCol-1)
					if !isTopCorner && !isBottomCorner {
						matrix[r][c] = ' '
					}
				}
			}

			for i, text := range msgLines {
				text = strings.TrimSpace(text)
				row := startY + i
				startC := centerX - (len(text) / 2)
				for j, char := range text {
					col := startC + j
					if row < len(matrix) && col < len(matrix[row]) {
						matrix[row][col] = char
					}
				}
			}

			var out strings.Builder
			for _, row := range matrix {
				out.WriteString(string(row) + "\n")
			}
			fmt.Print(out.String())

			time.Sleep(1800 * time.Millisecond)
		}
	}
}
