# asg 1

* Antoine Grondin, 6276497:
    * implemented the `document`/`index`/`query`/`ranking`/`similarity`/`term`
packages.
* Raj Pathak, 5650066:
    * implemented the `preprocess` package.
    * integrated the `bleve`-based library.
    * wrote this document.

## Functionality

We have two implementations of the IR system.

* A completely fresh implementation, in the subpackages.
* An implementation using off-the-shelf software, using the `bleve` search engine.

## Algorithms, data structures and optimizations:

Our implementation uses:

- Preprocessing:
    - A tokenizer that keeps only words or hashtags.
    - A stopword list of ~100 words.
    - A Porter2 stemmer for English.
    - A term booster for twitter hashtags (5x boost).
- Indexing:
    - Mapping from term to document is done with an in-memory hashtable, for O(1)
    term-to-document lookup.
    - An array of all documents is kept in the index, for fast document iteration.
    - Documents hold terms in a bag, implemented using an in-memory hashset, which
    keeps track of their frequency.
- Searching:
    - Weigthing is done using `TFIDF`.
    - Scoring is computed using cosine distance.

The index and documents are held in memory.

### Optimizations

TFIDF are cached once computed, so that future lookups be faster.

## How to run

Precompiled binaries are included. Otherwise, you can compile your own.

### Compile

You will need to have Go 1.4.1 installed to compile the program.

- See: https://golang.org/doc/install
- OS X installer: https://storage.googleapis.com/golang/go1.4.1.darwin-amd64-osx10.8.pkg
- Linux archive: https://storage.googleapis.com/golang/go1.4.1.linux-amd64.tar.gz
- Windows installer: https://storage.googleapis.com/golang/go1.4.1.windows-amd64.msi

Once Go is installed, from the root of the project, compile and run the program:

```bash
# compile it
go build -o trec_score ./cmd/bleve_reference/
# run it
./trec_score -corpus="twitter_entries.txt" -trec="trec_top_queries.xml" -output="result.txt"
```

### On OS X:

From a terminal (Terminal.app) at the root of the project:

```bash
$ ./trec_score_os_x -corpus="twitter_entries.txt" -trec="trec_top_queries.xml" -output="result.txt"
```


### On Linux:

From a terminal at the root of the project:

```bash
$ ./trec_score_linux -corpus="twitter_entries.txt" -trec="trec_top_queries.xml" -output="result.txt"
```

### On Windows:

From a `cmd` prompt, at the root of the project:

```bash
$ ./trec_score_win.exe -corpus="twitter_entries.txt" -trec="trec_top_queries.xml" -output="result.txt"
```

## FAQ

### Discuss your final results.

Final results are quite good.

```
num_q           all 49
num_ret         all 39387
num_rel         all 2640
num_rel_ret     all 2238
map             all 0.2650
gm_ap           all 0.1982
R-prec          all 0.2997
bpref           all 0.2787
recip_rank      all 0.5941
ircl_prn.0.00   all 0.6777
ircl_prn.0.10   all 0.5132
ircl_prn.0.20   all 0.4235
ircl_prn.0.30   all 0.3524
ircl_prn.0.40   all 0.3249
ircl_prn.0.50   all 0.2863
ircl_prn.0.60   all 0.2056
ircl_prn.0.70   all 0.1703
ircl_prn.0.80   all 0.1342
ircl_prn.0.90   all 0.0913
ircl_prn.1.00   all 0.0192
P5              all 0.4000
P10             all 0.3408
P15             all 0.3116
P20             all 0.2959
P30             all 0.2823
P100            all 0.1845
P200            all 0.1295
P500            all 0.0781
P1000           all 0.0457
```

Our recall and precision numbers, based off the data, were the following:

```
recall = 2238 / 2640 = 0.8477
precision = 2238 / 39387 = 0.05682
```

![document-level-vs-precision-image][document]
![recall-level-vs-rank-image][recall]

Here we can directly observe the phenomenon of the trade-off between recall and precision, as discussed in class. Looking at both the incremental recall level and document number level precision averages, we can see this taking place as the precision drops very dramtically upon approaching full recall.

Based on the R precision, we can see that most of the relevant documents are not actually among the first recalled, probably due to the overall very high recall value.

However, looking at the average precision values, we see the system has a 20-27% rating per query, leaving it overall to be satisfactory.

[document]: https://cloud.githubusercontent.com/assets/1189716/6180746/3651e782-b2f7-11e4-8118-7f4b6a7cf7eb.jpg "Document Level vs Precision"
[recall]: https://cloud.githubusercontent.com/assets/1189716/6180749/38481336-b2f7-11e4-9133-60b9f7d87654.jpg "Recall Level vs Precision"

### How big was the vocabulary?

86764 terms in 45899 tweets.

### Include the first 10 answers to queries 1 and 25.

- `BBC World Service staff cuts`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 4.108 | `30260724248870912` | `BBC World Service outlines cuts to staff http://bbc.in/f8hYAT`
| 3.220 | `30198105513140224` | `BBC News - BBC World Service cuts to be outlined to staff http://www.bbc.co.uk/news/entertainment-arts-12283356`
| 2.298 | `29983478363717633` | `BBCNews] Major cuts to BBC World Service: BBC World Service is to close five of its language services, with th... http://bbc.in/e2vlpX`
| 2.254 | `29993695927336960` | `Major cuts to BBC World Service: BBC World Service is to close five of its language services, with the likely lo... http://bbc.in/eftjNe`
| 2.215 | `33823403328671744` | `World Service Cuts: Why We Need the BBC http://bit.ly/fhLqaS`
| 2.133 | `30275282464153600` | `BBC World Service to cut [...] a quarter of its staff - after losing millions in funding from the Foreign Office. http://bbc.in/hyGSHi`
| 2.133 | `30016851715031040` | `A statement on the BBC World Service, ahead of staff briefings/ further details on Weds http://bbc.in/dFfXIW #bbcworldservice #bbccuts`
| 2.112 | `30236884051435520` | `BBC confirm World Service cuts http://f.ast.ly/Q3Rfa`
| 2.112 | `30216589932503040` | `BBC World Service cuts = global audience to fall by 30 million.`
| 2.020 | `30167063326629888` | `Quarter of BBC world service staff to go, uk foreign office grant reduction of 17.5%.`
- `2022 FIFA soccer`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.035 | `35048150574039040` | `who cares about the child marriage...the 2022 #soccer world cup will be held in #Qatar anyways :(( #secclinton #fifa`
| 1.577 | `32242968895168512` | `Qatar's 2022 FIFA World Cup Stadiums: http://wp.me/p18Mk5-2Js`
| 1.460 | `33627853044056064` | `Lovin this day off of doing nothing...FIFA soccer all day! #goooooaaaallllll`
| 1.445 | `29058771531595776` | `Sports digest: FIFA discusses moving 2022 World Cup in Qatar from summer to winter - San Jose... http://bit.ly/e9y8Ep #fifa`
| 1.376 | `34063614319136769` | `Jogando FIFA Soccer 11. http://raptr.com/DonDiegoBRA`
| 1.376 | `33562935863410688` | `Playing FIFA Soccer 11. http://raptr.com/jfooks`
| 1.376 | `32870931382673409` | `Playing FIFA Soccer 10. http://raptr.com/Primexxx`
| 1.376 | `32445789523419136` | `is playing FIFA Soccer 11. http://raptr.com/Srg11`
| 1.376 | `32301707404771328` | `Playing FIFA Soccer 11. http://raptr.com/morebk`
| 1.376 | `32299263832621057` | `Playing FIFA Soccer 11. http://raptr.com/liam017`
- `Haiti Aristide return`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.400 | `29296574815272960` | `Haiti – Aristide : His return, an international affair… – Haitilibre.com http://bit.ly/gzyLXG #haiti`
| 3.284 | `32204788955357184` | `Haiti opens door for return of ex-president Aristide http://tf.to/fJDt`
| 3.144 | `29278582916251649` | `Haiti - Aristide : His return, an international affair... - http://haitilibre.com/fben.php?id=2193`
| 3.064 | `29613127372898304` | `If Duvalier Can Return to Haiti, Why Can’t Aristide? – New America Media http://bit.ly/eCWStk #haiti`
| 2.964 | `31861291236724738` | `Haitian Politics - Rev Jeremiah Wright Wants ARISTIDE To Return To Haiti: Yes, It will be good for Aristide to r... http://bit.ly/f2hFSi`
| 2.911 | `34694060157435904` | `Former Haitian President Aristide has been issued with a new passport enabling him to end exile and return to Haiti, from AFP`
| 2.911 | `34692276609351680` | `Former Haitian President Aristide has been issued with a new passport enabling him to end exile and return to Haiti, from AFP`
| 2.903 | `32383831071793152` | `Yah Haiti: Haiti allows ex-president's return: Jean-Bertrand Aristide, who was Haiti's first democratically ele.... http://bit.ly/hLAgwO`
| 2.812 | `32443364628500480` | `#MIAMI Haiti to issue ex-president Aristide with passport, clearing way for him to return http://bit.ly/fAV5fB`
| 2.812 | `32411439918489600` | `Haiti allows ex-president's return: Jean-Bertrand Aristide, who was Haiti's first democratically elected leader,... http://aje.me/fQ4j4T`
- `Mexico drug war`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.921 | `29400624374222849` | `Mexico's Drug War on CNBC. #sad`
| 3.128 | `30027043655655424` | `Clinton backs Mexico's drug war http://bit.ly/ehlzQp`
| 3.128 | `29684273590042624` | `Clinton supports Mexico in 'messy' drug war http://wapo.st/guFtd1`
| 2.995 | `32851298193768448` | `LPGA cancels Mexico stop for drug war concerns - http://es.pn/fdduAA`
| 2.995 | `30470121625485312` | `Legalization of drugs would end the drug war and related violence in Mexico http://fb.me/tXGnLjw3`
| 2.679 | `29878816281206784` | `Why Hillary Clinton flagged judicial reform as 'essential' to Mexico's drug war http://su.pr/31zNhB`
| 2.621 | `29756448648994816` | `Headline News - Clinton backs Mexico's drug war: US Secretary of State Hillary Clinton, on a visit to Mexico, sa... http://bbc.in/g1QFYa`
| 2.445 | `29903758779490304` | `Why Hillary Clinton flagged judicial reform as 'essential' to Mexico's drug war: ... for years to overhaul inade... http://bit.ly/ifscaU`
| 2.445 | `29887625179435009` | `Why Hillary Clinton flagged judicial reform as 'essential' to Mexico's drug war #Central #America http://myfeedme.com/m/11120095`
| 2.320 | `30306064587030528` | `The Call: 2011's Top Risks: In Mexico, the drug war heats up: More broadly, the political consen... http://bit.ly/e3XUG2 #tcot #tlot #p2`
- `NIST computer security`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.702 | `29644789049724928` | `NIST Addresses a Security Threat that Challenges Most Information Security Programs http://budurl.com/DSnist #B2B`
| 1.663 | `33288536677421057` | `New NIST Guidance Tackles Public Cloud Security http://ow.ly/3PUFB`
| 1.663 | `33188770152976385` | `New NIST Guidance Tackles Public Cloud Security http://bit.ly/ezdDcu`
| 1.501 | `33579862476197888` | `Security News - NIST Issues Cloud Security Guidelines: The government standards body has launched a wiki to get ... http://bit.ly/e4miOd`
| 1.487 | `33158966397763585` | `New NIST Guidance Tackles Public Cloud Security - BankInfoSecurity.com http://goo.gl/fb/o8aaF`
| 1.398 | `30566466117967872` | `Computer Security: Security Devices, Computer Lock, Laptop ...: It's time to install computer security gadget fo... http://bit.ly/dKKwyg`
| 1.397 | `33581589627666432` | `NIST Issues Cloud Security Guidelines: The government standards body has launched a wiki to get feedback on its ... http://bit.ly/icedHU`
| 1.358 | `32977594907492352` | `New NIST Guidance Tackles Public Cloud Security: 2 Other Special Pubs on Cloud Defs, Virtualization. http://bit.ly/f8Vgqi`
| 1.112 | `32670074020044800` | `The Computer Attacks You've Never Heard Of [Security] http://jfish.me/hBqgx0`
| 1.112 | `30774950642057217` | `A guide to computer network security http://bit.ly/8Vgt94 #learning`
- `NSA`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.640 | `29777900525850624` | `nsa made me stupid.....`
| 2.752 | `33752832922353664` | `@mapet0819 nsa showtime po kau today?`
| 2.752 | `32531206918635520` | `Tomorrow, NSA National Business Park Expo!`
| 2.752 | `31944622758428672` | `@ShakurDC_NSA YAAA LETS GOOO SHAKKERRR`
| 2.574 | `35005178885181441` | `Get the latest NSA job opportunities delivered directly to your phone through the NSA Career Links app! Available for iPhone and Android.`
| 2.574 | `31098028576215040` | `@joia24_7 awww i wanted to watch nsa!! =[`
| 2.574 | `30039680401539072` | `@sportsguy33 if i forwarded this to my contacts at the NSA they would pay you a special visit:-):-):-)`
| 2.574 | `29816595014483968` | `NSA - Security on Gizmodo- http://bit.ly/dEGVPU`
| 2.574 | `29363446097117184` | `nsa o meu braço direito começou a doer :/`
| 2.427 | `34713303489978368` | `Postando as informações da #RuaDasBarracas nsa comunidades referentes a Pinhal.`
- `Pakistan diplomat arrest murder`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.021 | `30790044126027776` | `Update: U.S. diplomat charged with Pakistan double murder - BBC http://bbc.in/hFObgw`
| 1.881 | `30723400687165441` | `Uh oh, American diplomat charged with double murder in Pakistan: http://yhoo.it/fG04tY He should have just called in a drone.`
| 1.671 | `34857364532236288` | `DTN India: US mounts pressure on Pakistan to release 'illegally detained' murder-accused diplomat: Islamabad, Fe... http://bit.ly/gRaUKb`
| 1.671 | `30809856231342080` | `BBC News - US diplomat charged with Pakistan double murder http://www.bbc.co.uk/news/world-south-asia-12304457`
| 1.162 | `33353153395040257` | `Case Of Jailed Diplomat In Pakistan Fuels Anger http://interesting.rssnewest.com/case-of-jailed-diplomat-in-pakistan-fuels-anger/`
| 1.130 | `33094359847014401` | `Pakistan court extends detention of US diplomat - msnbc.com`
| 1.022 | `32503347730714624` | `Pakistan court blocks handing over of U.S. "diplomat" *http://is.gd/nCF2BE`
| 1.022 | `31647814786228224` | `U.S. Demands Release of Diplomat in Pakistan http://buz.tw/pcVlo`
| 1.022 | `31413608780931072` | `NDTV: US asks Pakistan to release diplomat http://snipurl.com/1xerb8`
| 1.019 | `31329744452591617` | `RT @TheNewsBlotter: ISLAMABAD, Pakistan (AP) -- US demands release of diplomat in Pakistan http://apne.ws/htHOfL`
- `phone hacking British politicians`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.770 | `34467845525995520` | `Every Famous British Person to Sue Tabloid Over Phone Hacking http://bit.ly/fIQEME`
| 1.587 | `29281084667596800` | `Brown may be British phone-hacking target: LONDON, Jan. 23 (UPI) -- Former British Prime Minister Gordon Brown ... http://bit.ly/gcER5e`
| 1.066 | `30300450339164160` | `Ian Edmondson sacked from NOTW over phone hacking...`
| 0.997 | `30303893518819329` | `Met in fresh phone hacking probe: Police launch a fresh investigation into phone hacking after receiving "significant new information...`
| 0.996 | `29192127565004803` | `To Spy Politicians, British Aide to Prime Minister Resigns: http://cot.ag/hLRsJM`
| 0.948 | `30429664170213377` | `British Tabloid Dismisses Editor Over Hacking Scandal http://nyti.ms/hZk0OH`
| 0.940 | `32480508667494400` | `Go to my blog in a few minutes for an interesting development in phone hacking scandal`
| 0.940 | `32473801157517312` | `Go to my blog in a few minutes for an interesting development in phone hacking scandal`
| 0.940 | `31132093111074816` | `@SallyBercow @TomHarrisMP I talk more on Twitter than on the phone,no hacking necessary lol`
| 0.940 | `30540202019655680` | `“There are questions about why the investigation [into phone hacking] was stopped when it was”- Jenny Jones #r4today`
- `Toyota Recall`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.214 | `31987981896126464` | `lexus also included in the latest Toyota recall`
| 3.214 | `30225672072855552` | `Toyota to recall 1.4 million vehicles -`
| 2.835 | `30282297689251840` | `Toyota Recall http://j.mp/f2rCNQ via @AddToAny`
| 2.835 | `30381116489736193` | `Toyota Recall 2011: 245,000 Lexus Vehicles Recalled - Associated ...: Toyota has issued a recall for a num... http://tinyurl.com/4ux7hac`
| 2.689 | `30459074709557248` | `New Toyota Recall Involves Lexus Models - http://newzfor.me/?cotf`
| 2.689 | `30439920199798787` | `New Toyota Recall Involves Lexus Models - http://newzfor.me/?cotf`
| 2.689 | `30203004422463488` | `News from Ireland. Fresh recall setback for Toyota: Toyota has said it will recall 1.7 million ... http://tinyurl.com/67x3ugz Irish News`
| 2.689 | `30151108915625984` | `DTN Financial News: Toyota to recall 1.7 million vehicles worldwide (AFP): AFP - Toyota Motor is to recall ... http://bit.ly/iiLzpT`
| 2.613 | `30325626942529536` | `Toyota's latest recall - more mechanical misery or should they be applauded? Plus, have you suffered a #car recall? http://bit.ly/g9UYuN`
| 2.564 | `35067946019590144` | `Feds clear Toyota of electronic causes in recall probe - http://newzfor.me/?cyvr`
- `Egyptian protesters attack museum`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.785 | `31075323059638272` | `Protesters forming teams to protect the Egyptian Museum from thieves. #Egypt #Jan25 #SidiBouzid`
| 1.709 | `31245512300568577` | `Confirmed: Egyptian protesters & activists successfully protected the national museum from looters #Egypt #Jan25`
| 1.528 | `31077260689674240` | `Protesters form human shield around Egyptian National Museum. Risking their lives to save their history. #Jan25 #Egypt`
| 1.495 | `32275485807353856` | `@qronoz: Two ancient mummified heads lie on the floor of the Egyptian Museum after the weekend attack by looters. http://bit.ly/dG8Et1`
| 1.395 | `31855298419359745` | `Anti-Government Protesters in Cairo Smash Treasures and Mummies in Egyptian Museum: CAIRO (REUTERS).- Looters br... http://bit.ly/hrKiH7`
| 1.358 | `31609613321244672` | `Egyptian protesters ransack Cairo museum, smash mummies: The mass anti-government protests in Egypt took a toll ... http://bit.ly/gFhn9Q`
| 1.090 | `31518116253007874` | `Rosicrucian Egyptian Museum: Rosicrucian Egyptian Museum http://bit.ly/fZhNnq`
| 1.044 | `31401050451738624` | `Horrified at scenes of looting of the Egyptian Museum`
| 1.033 | `31499543946207232` | `A collection of tweets on the the Egyptian Museum in Cairo: Is the Egyptian Museum Under Threat? http://bit.ly/eOgAHp #Jan25 #egypt #museum`
| 0.933 | `31076061462659072` | `al Jazeera: Thousands of Egyptian youth form human shields to protect the Egyptian Museum. #Museum #Egypt #Jan25`
- `Kubica crash`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.789 | `34199299428581376` | `F1 star Kubica involved in serious crash - Yahoo! Eurosport http://yhoo.it/ePoLOB`
| 2.628 | `34530871604289536` | `I'm not a #Kubica fan but this is sad - Robert Kubica's F1 career at risk after rally crash http://www.bbc.co.uk/news/world-europe-12378835`
| 1.372 | `30633225000849408` | `To all Robert Kubica fan out there, please do follow @F1Ellen she is cool and awesome kubica fan too :)`
| 1.216 | `33281011269902336` | `Kubica fastest in Valencia? I knew the renault would be epic!`
| 1.140 | `30690569252507648` | `Crash, crash, burn.`
| 1.140 | `33289317065433088` | `crash into me.`
| 1.140 | `32406655899533312` | `Crash **`
| 1.140 | `31976428450877440` | `About To CRASH!`
| 1.140 | `31587605195792384` | `#crash`
| 1.140 | `31561746011398144` | `Crash on`
- `Assange Nobel peace nomination`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.361 | `32117461922877440` | `Nobel Peace Prize nomination for WikiLeaks founder? http://bit.ly/gKQ1Kl`
| 1.305 | `31940748085563392` | `维基解密的创始人阿桑奇可能会参选2011诺贝尔和平奖：http://www.rnw.nl/english/bulletin/nobel-peace-prize-nomination-wikileaks-founder`
| 1.137 | `29819392954994688` | `#UnlikelyHeadlines Osama wins Nobel peace prize`
| 1.063 | `32181966761631744` | `Nobel Peace prize for WikiLeaks? http://is.gd/ykEIRU`
| 1.063 | `32112594718294016` | `Nobel Peace prize for WikiLeaks? http://is.gd/ykEIRU`
| 1.063 | `31939614583291904` | `Nobel Peace prize for WikiLeaks? http://is.gd/ykEIRU`
| 1.063 | `31939483347722242` | `Nobel Peace prize for WikiLeaks? http://is.gd/ykEIRU`
| 1.002 | `29948475433029633` | `#unlikelyheadlines GEORGE BUSH WINS NOBEL PEACE PRIZE! Ha`
| 0.997 | `29939186026946560` | `NEWS: The Nomination for Prince George Peace River Has Begun http://ow.ly/1b1ptC`
| 0.976 | `29210686835916800` | `Irony: The 2009 Nobel Peace Prize winner hosted a State Dinner for the man who is holding the 2010 Nobel Prize Peace Prize winner in prison.`
- `Oprah Winfrey half-sister`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 4.114 | `29561670661570560` | `Oprah Winfrey announces that she has a half-sister who she never knew existed. She'll announce which half on an upcoming show.`
| 3.445 | `29560222905278464` | `who gives a shit - RT @CNNshowbiz: Oprah Winfrey announces that she has a half-sister who she never knew existed http://on.cnn.com/dL7MNS`
| 2.850 | `29564454236594176` | `I am Oprah's half sister.`
| 2.501 | `30345378041692161` | `The Clicker - Oprah's dad says he learned sister secret when he watched the show: When Oprah Winfrey announced h... http://bit.ly/icUZGp`
| 2.395 | `29579396469760000` | `Oprah Winfrey finds sister she didn't know she had! She's known since 2007?! http://ow.ly/3JbDB`
| 2.253 | `29558797357809664` | `Who hasn't dreamed of being Oprah's half-sister?`
| 2.023 | `29540081047961602` | `Oprah Winfrey: What is Her Family Secret?: By Lyneka Little Oprah Winfrey's Family Secret: Oprah Winfrey says th... http://bit.ly/h7Jrgu`
| 2.011 | `29557749993963522` | `Does Oprah have a half-sister that she didn't know existed? http://bit.ly/fjDxat #Oprah`
| 1.974 | `29551669075251200` | `Oprah Winfrey Big Family Secret - Secret Show oprah winfrey http://ping.fm/v3jot`
| 1.921 | `29563563311894528` | `@Oprah ruined my fun already. She has a secret half-sister. NEW GAME! #doesstedmanexist`
- `release of "The Rite"`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 0.932 | `31529870823526400` | `Is that rite?`
| 0.932 | `30955134796177408` | `Rite`
| 0.761 | `30730642933153792` | `@LadyJay85 rite rite catch u then :)`
| 0.761 | `32683230717022208` | `@cocoShanaiel rite`
| 0.761 | `32627227011055616` | `@TeamMinajMuny. :/ rite here.`
| 0.761 | `32508295373651968` | `@MissStormrose or rite then`
| 0.761 | `32194856520519681` | `Yea she rite`
| 0.761 | `32192029555429376` | `sumthing not rite........`
| 0.761 | `31938455734853632` | `@1LoveMaria I no rite.`
| 0.761 | `31916819551879168` | `Cinema..."The rite"`
- `Thorpe return in 2012 Olympics`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.821 | `32665910925852672` | `Thorpe ends retirement, looks to 2012 Olympics (The Arizona Republic) http://feedzil.la/huBOUo`
| 1.803 | `32678457083170816` | `Australia's Ian Thorpe to return for London Olympics http://bit.ly/eZW5B6`
| 1.755 | `33276913967435776` | `Ian Thorpe is returning to contest the 2012 London Olympics. Picture: AFP ... http://bit.ly/hJVceZ`
| 1.695 | `32674038430040064` | `Just heard on the radio that Ian Thorpe is back in the swimming! He's going to have some serious Dutch rivals @ Olympics 2012 :-))`
| 1.642 | `32703076989140992` | `Heard Ian Thorpe is making a comeback for 2012 Olympics. Now rumours of further comebacks. AJ Boyse unavailable for comment at this time.`
| 1.537 | `32641649188274176` | `Please reconsider, Ian Thorpe http://www.foxsports.com.au/other-sports/ian-thorpe/healtfelt-advice-for-ian-thorpe-please-reconsider-your-swimming-return-for-2012-london-olympic-games/story-fn7rv3iq-1225998678900 …`
| 1.515 | `33087998794932224` | `Morning.Thorpes return can only be good for the sport and the Olympics.What a match up, Thorpe v phelps, u now it will be race of the games.`
| 1.433 | `32760806558928896` | `NYT Sports Sports Briefing | Swimming: Thorpe Ends Retirement, Looks to 2012 Olympics: Five-time Olympic gold me... http://bit.ly/gY5JPY`
| 0.957 | `32386030434787328` | `Mixed reactions to Thorpe return http://bit.ly/flRn4f`
| 0.945 | `32608279829938177` | `He likes the training and he likes the racing. Glad to be able to go to another Olympics before he is 30. #Thorpe`
- `release of "Known and Unknown"`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 0.700 | `29585186899365888` | `Donald Rumsfeld Reflects on Writing of Memoir: When Rumsfeld began thinking about his book, "Known and Unknown,"... http://bit.ly/eVnibq`
| 0.516 | `31239257569894401` | `RELEASE RELEASE RELEASE !! I WILL BUYY !!! :D`
| 0.498 | `32143203054395392` | `#np release me, release my bodyyy`
| 0.455 | `30607693014110208` | `Release to refresh`
| 0.455 | `30469928335183872` | `Release the Validation!`
| 0.455 | `30447130980122624` | `Release The Kraken!!!`
| 0.455 | `30353344560103425` | `#TYPO3LTS we have a Release!`
| 0.455 | `30232523275505664` | `Early release`
| 0.394 | `32773048323014656` | `Release your inner Groundhog`
| 0.394 | `32301131518447616` | `So we have early release wednesday?`
- `White Stripes breakup`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.116 | `32871838174416897` | `The White Stripes announce breakup http://on.today.com/eEdSvJ /say it ain't so!`
| 2.884 | `32867803606290433` | `More on the White Stripes' breakup: they want to "preserve what is special about the band & have it stay that way" http://bit.ly/hPeZDQ`
| 2.544 | `32876528131899392` | `NEWS+VIDEOS | The White Stripes Call It Quits | http://bit.ly/eHg3LH (via @pitchforkmedia) #Read #Breakup #RIP #Bummer`
| 2.213 | `32871790007033857` | `White Stripes! NO!!!!!!!!!!!!!!`
| 1.917 | `32876849235230720` | `White Stripes - äntligen!`
| 1.917 | `32874195033522176` | `RIP White Stripes.`
| 1.917 | `32871097816842240` | `Something about the White Stripes.`
| 1.917 | `32861990636494848` | `WHITE STRIPES NOOOOOO!!!!!`
| 1.714 | `32869378253012992` | `A shame that the white stripes have broken up.`
| 1.714 | `32867603496046592` | `R.I.P White Stripes.`
- `William and Kate fax save-the-date`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.054 | `30273593405345792` | `Prince William And Kate Middleton's Wedding: Did You Get Your Save-The-Date Fax? http://bit.ly/hs1nAm`
| 2.980 | `31396010894819329` | `William and Kate fax their save-the-dates - http://newzfor.me/?cvig`
| 1.405 | `30188073790742528` | `Prince William and Kate Middleton's wedding guests invited by fax - Telegraph.co.uk http://goo.gl/fb/2s7nw`
| 0.813 | `29964583695294464` | `Prince William and Kate Middleton to Invite Sarah Ferguson to ...: Prince William and Kate Middleton to Invite S... http://bit.ly/hqVm9x`
| 0.797 | `31014649839226880` | `Royal Roundup: Bald Spots and Fake Kates!: While we eagerly await our save-the-date fax for the royal wedding, w... http://bit.ly/gWJxNx`
| 0.772 | `29539195265490944` | `'William en Kate-effect' http://telegraaf.nl/s/8821683`
| 0.729 | `29248037750571008` | `Kate Middleton | Prince William Royal Outcast Sarah Ferguson calls his marriage to Kate … http://dlvr.it/Dwv4k`
| 0.704 | `29579118953627648` | `Where do you think Prince William and Kate Middleton should spend their honeymoon? Let us know your ideas!`
| 0.704 | `29471305610825728` | `And the BBC's coverage of William and beautiful/down-to-earth/magnetic Kate continues. Today it was about plates...`
| 0.704 | `29356433623490560` | `OMG William & Kate are getting married on my BIRTHDAY this year!! Wow!! I should be dubbed royalty :O`
- `Cuomo budget cuts`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.401 | `32579076275306496` | `Cuomo Cuts Spending in Budget Plan http://bit.ly/dHVQla`
| 2.778 | `32550327760723968` | `Cuomo’s Budget Cuts Spending on Schools and Medicaid: Gov. Andrew M. Cuomo on Tuesday proposed a $132.9 billion... http://nyti.ms/dGLeNP`
| 2.777 | `34046080261824512` | `#nano_bio Nanotechnology gets full funding in Cuomo budget, amid cuts http://ow.ly/1bcnOp`
| 2.777 | `32324277558575104` | `Cuomo budget cuts spending for first time since mid-199... - http://fwix.com/a/23_1be98382c2`
| 2.714 | `32550210244706304` | `Cuomo’s Budget Cuts Spending on Schools and Medicaid: Gov. Andrew M. Cuomo on Tuesday proposed a $132.9 bil... http://nyti.ms/gkUOAO nyt`
| 2.599 | `32403820008968192` | `Cuomo exposes dirty budget trick - $10B cuts announced today http://www.nypost.com/p/news/local/cuomo_exposes_dirty_trick_DRgwu2QVC7WnpgORXaWZzM … via @newyorkpost`
| 2.468 | `32864634478272512` | `lakesuccessny: lakesuccessny: Cuomo's NY Budget Calls For Spending Cuts, Layoffs – CBS News |… http://goo.gl/fb/niLSD`
| 2.405 | `32342810979999744` | `Why spending's always up: The 2011-12 budget Gov. Cuomo unveils today will include substantial cuts in hundreds ... http://nyp.st/hejdVT`
| 1.997 | `33587838331129856` | `#scariestwordsever budget cuts`
| 1.831 | `32503415393226752` | `Cuomo proposes deep cuts to SUNY and CUNY. Shameful!`
- `Taco Bell filling lawsuit`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.306 | `29853985930219520` | `That ain't necessarily "beef" in your Taco Bell burrito...a new lawsuit wants the chain to label it "taco meat filling" instead.`
| 3.159 | `31082136219947008` | `Taco Bell Counters 'Meat Filling' Charges in Lawsuit With Print, Web Effort http://goo.gl/fb/H9wcB`
| 3.106 | `29906116062220290` | `Lawsuit: Taco Bell Ground Beef Is Really Just "Meat Filling" - @consumerist http://consumerist.com/2011/01/lawsuit-says-taco-bell-ground-beef-is-really-just-taco-meat-filling.html?utm_source=streamsend&utm_medium=email&utm_content=13297631&utm_campaign=Fo`
| 2.389 | `29913837511647232` | `OMG Taco Bell has been using "taco filling" not meat for decades...this is not new people`
| 2.164 | `31161931205181440` | `Eating meat filling all 35% of it (@ Taco Bell) http://4sq.com/fSf4Es`
| 2.142 | `32218912527482880` | `Lawsuit to Taco Bell: Where?s the Beef? http://daily.rssnewest.com/lawsuit-to-taco-bell-wheres-the-beef/`
| 2.118 | `30075676803465217` | `Hmm... I really want to go to Taco Bell and ask for taco meat filling now...Thank you #NPR`
| 2.099 | `32157053631860736` | `Only people who think of Taco Bell as health food.RT @GourmetFury: Do you really give a crap what goes into Taco Bell's filling?`
| 2.072 | `30400693814697985` | `Bored critics have a beef with Taco Bell 19s meat filling - http://newzfor.me/?c4qf`
| 2.046 | `30283063699177472` | `Oh, my... Taco Bell's "Taco Meat Filling" is only 36% beef... http://gizmodo.com/5742413/ #fb`
- `Emanuel residency court rulings`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.597 | `29963031203020800` | `BREAKING: Illinois Supreme Court just ruled that Emanuel's name must be included on ballot until residency issue is resolved.`
| 1.597 | `29610928903299072` | `The Chicago Sun-Times reports that the appeals court has ruled Emanuel does not meet the residency require... http://tinyurl.com/4v93ptb`
| 1.597 | `29606637102702593` | `Court rules that Rahm Emanuel should be removed from Chicago mayoral ballot, does not meet residency requirements http://bit.ly/fsG9so`
| 1.549 | `29985282845581312` | `RT @USATODAY: Court puts Rahm Emanuel back on ballot, agrees to hear residency case http://usat.ly/dO6ePn`
| 1.428 | `29612419269525504` | `Appeals court says Rahm Emanuel can't be on Chicago mayoral ballot because he doesn't meet residency rules. http://apne.ws/fYxbkA`
| 1.324 | `29613409045577729` | `@governorrod What are your thoughts on the Emanuel residency ruling?`
| 1.170 | `33190142852206592` | `Search Tennessee Court Rulings in Google Scholar http://goo.gl/fb/9Odj5 #tcot #teaparty`
| 0.978 | `29832102534971392` | `Rahm Emanuel Residency Twitter Reactions (PHOTOS) http://bit.ly/fa0HJz`
| 0.917 | `29627818857996289` | `BREAKING: Appellate Court Rules against Rahm residency. http://bit.ly/gPw2Xx`
| 0.894 | `30764955837927424` | `IL Supreme Court overturns Appellate Court and rules #Rahm meets residency requirement. Rahm stays in the race.`
- `healthcare law unconstitutional`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.758 | `32168717513138176` | `#BreakingNews: Florida Judge Rules #HealthCare Law Is Unconstitutional`
| 3.758 | `32167894922043392` | `#BreakingNews: Florida Judge Rules #HealthCare Law Is Unconstitutional`
| 3.758 | `32167461910478848` | `#BreakingNews: Florida Judge Rules #HealthCare Law Is Unconstitutional`
| 3.362 | `32171528254660608` | `RT@foxnews: #BreakingNews: Florida Judge Rules #HealthCare Law Is Unconstitutional`
| 3.205 | `32275892562567168` | `Federal judge says healthcare law unconstitutional http://tinyurl.com/4umv9nh`
| 3.069 | `32170314901233664` | `Florida judge declares parts of healthcare law unconstitutional in suit involving SC and other states`
| 2.948 | `32218747884277760` | `Healthcare law unconstitutional and therefore void...hmmm where have I heard that before? Oh, that's right... ME!! #fb #mamagrizzly`
| 2.948 | `32175018792325120` | `Well this a surprise and not a good thing-Healthcare law declared unconstitutional: http://bit.ly/icZyWN`
| 2.841 | `32169956334370816` | `BREAKING NEWS: Federal judge in Florida strikes down parts of the Obama administration’s healthcare law as unconstitutional.`
| 2.745 | `32173003508813824` | `RT @RichardA Here we go: federal judge in Florida has ruled provisions of the healthcare law unconstitutional || #hcr #singlepayer`
- `Amtrak train service`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.402 | `33310482420998145` | `Ugh my #amtrak train is delayed again.`
| 1.713 | `30599622661505024` | `AMTRAK service north of NYC now cancelled @toddstarnes`
| 1.699 | `34729493578907648` | `Derailed Amtrak train a real headache for LIRR http://huff.to/hGwo0e`
| 1.641 | `32638038530072576` | `6 train service has resumed regular service at this time.`
| 1.620 | `35047250950361089` | `Amtrak train derails after leaving New York station - http://newzfor.me/?csjr`
| 1.620 | `34080830133379072` | `Amtrak train strikes eagle in rare encounter: http://wapo.st/eKc951`
| 1.620 | `33476555191615488` | `Amtrak Train Murders Bald Eagle [Felonies] http://gaw.kr/ecpp8m`
| 1.551 | `29626342530088960` | `New Oil: Sunday Train: Going on the Attack for Amtrak http://bit.ly/gjAPIq`
| 1.519 | `33011776777883648` | `Amtrak sues Detroit over fire truck, train collision - VIDEO - http://www.firerescue1.com/apparatus/articles/966367-Amtrak-sues-Detroit-over-fire-truck-train-collision/ … via @AddThis`
| 1.490 | `34744935773118464` | `Amtrak train derails after leaving New York station http://bit.ly/f9J2k7 #news`
- `Super Bowl, seats`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.055 | `34646404349562880` | `400 denied seats at Super Bowl: Some 400 ticket-holders were denied seats at the Super Bowl Sunday because certa... http://bit.ly/fRJkf3`
| 2.718 | `34518627260567552` | `Super Bowl fans sent home because seats unsafe http://bit.ly/fwvuwN`
| 2.718 | `34459879930470400` | `Super Bowl: Celebrities star in the seats, not just on the field http://bit.ly/ev0F5l`
| 2.718 | `34416067287715840` | `Who were all Thes people in the parking lot? #super bowl -people that got kicked out bad seats?`
| 2.686 | `34631454545678336` | `Super Bowl fans denied seats: Cowboys Stadium apparently was not ready for the Super Bowl. Before the game, ab... http://wapo.st/igEA0p`
| 2.614 | `34495093771747328` | `Super Bowl Fans Sent Home Because Seats Unsafe: Cowboys Stadium wasn't ready for the Super Bowl. Before the... http://on.wesh.com/i8uoCB`
| 2.612 | `34734420598464513` | `Super Bowl Attendance Record Shot Thanks to Unfinished Seats http://goo.gl/Az5qU`
| 2.612 | `34389740572647424` | `Fans denied access to seats at Super Bowl - Yahoo! Sports http://ff.im/-xt5ZQ`
| 2.612 | `30322463783002112` | `Super Bowl Report fm Dallas: very few tickets left and all start above $1000 (4 worst seats)!`
| 2.601 | `34374988265947136` | `1,000 fans WITH tix to Super Bowl have no seats! Fire Marshall says the temp seats are not safe. They get triple face value refunds.`
- `TSA airport screening`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.040 | `31286354960715777` | `TSA shuts door on private airport screening program (CNN) http://feedzil.la/gD1tt6`
| 2.983 | `31738694356434944` | `TSA Shuts Door on Private Airport Screening Program – Patriot Update http://patriotupdate.com/2451/tsa-shuts-door-on-private-airport-screening-program?sms_ss=twitter&at_xt=4d45868911137f91,0 … via @AddThis`
| 2.830 | `31550836899323904` | `TSA shuts door on private airport screening program. Utter BS! - http://bit.ly/fx6Dgw #cnn`
| 2.337 | `32609015158542336` | `TF - Travel RT @Bitter_American TSA shuts door on private airport screening program - http://bit.ly/fx6Dgw #cnn:... http://bit.ly/eADg2G`
| 1.825 | `31320463862931456` | `TSA halts private screening program http://bit.ly/hUzJ3t`
| 1.825 | `30093525102108674` | `Obama makes fun of his own TSA screening procedures. Someone is off message! #sotu`
| 1.666 | `32685391781830656` | `Really looking forward to my TSA screening; haven't gotten laid in a couple of weeks.`
| 1.615 | `32541161675558912` | `TSA to Test New Screening at Hartsfield-Jackson: The TSA in coming days at Hartsfield-Jackson Atlanta Internatio... http://bit.ly/e8NW0S`
| 1.601 | `32528974961713152` | `Atl Business Chronicle: TSA to test new screening at Hartsfield-Jackson http://brkg.at/hR3W3y`
| 1.508 | `31773184512495616` | `The TSA has a unique ability to make an airport seem busy even when it isn't. #tsa`
- `US unemployment`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.224 | `30623689372340224` | `US: Requests for unemployment benefits up due to snow http://fxn.ws/gcshW1`
| 1.904 | `32546521291431937` | `The US Government is putting to DEATH the long-term unemployed, the 99ers, by STOPPING their paid-in-full unemployment insurance checks!!!`
| 1.844 | `33525722479861761` | `Missold Mortgage: US joblessness falls in January: Unemployment falls in January to 9% but the... http://bbc.in/eguXCV : Mortgage Claims`
| 1.844 | `30645832151728128` | `Foreclosure activity up across most US metro areas: The foreclosure crisis is getting worse as high unemployment and... http://dlvr.it/FCXVG`
| 1.789 | `33287991682138113` | `US initial jobless claims fall more-than-expected: Forex Pros – The number of people who filed for unemployment ... http://bit.ly/fywQwJ`
| 1.738 | `30982409809821698` | `REAL graduate unemployment figs actually far worse than suggested this week - 1000s of us are 'working' unpaid, http://bit.ly/gXX15x`
| 1.738 | `29263476715159553` | `For Obama, 'economy' is first priority, to generate more jobs in US: With the unemployment rate ending at 9.4 pe... http://bit.ly/ftVbrV`
| 1.692 | `31794789183651840` | `theresearchcenter.us.com - web info: 21 Jan 2011 TheResearchCenter $500 Unemployment Survey (theresearchcenter.u... http://bit.ly/fg3Ihd`
| 1.649 | `32912119968047104` | `US Still Home to 11.2M Illegal Immigrants -Deportations, unemployment doesn't diminish numbers -Newser http://bit.ly/e0lbKw #tcot`
| 1.649 | `32911735597826048` | `US Still Home to 11.2M Illegal Immigrants -Deportations, unemployment doesn't diminish numbers -Newser http://bit.ly/e0lbKw #tcot`
- `reduce energy consumption`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.805 | `32623902945447936` | `Night Curtains Reduce Energy Consumption And Help To Lower Your Energy Bill`
| 2.516 | `30642103809736704` | `Just joined Welectricity - a fun, FREE way for u & your friends to track & reduce energy consumption at home! Join me at http://bit.ly/EEmSN`
| 1.552 | `30726801869250560` | `What Tools and Tricks Do You Employ to Reduce Power Consumption? [Ask The Readers] http://bit.ly/gQr5An`
| 1.496 | `29859015223681024` | `One way to reduce our risk of cancer is to limit our consumption of meats, especially processed meats. /via @VeggieAdvisor`
| 1.482 | `30335672032169985` | `We improved the monitoring of our building systems and energy consumption to investigate where energy is wasted and how we can correct thes…`
| 1.466 | `29494814101733376` | `#Energy generation and consumption #infographics in USA http://bit.ly/fMV2GW`
| 1.399 | `29593306556014593` | `myworcester: Transport accounts for 75% of oil consumption, and oil is expected to peak by 2030 and then reduce ... http://bit.ly/i3IBsP`
| 1.397 | `33293927737987072` | `check out my SMPK at whywebpr.com/jenniferdorsey to find out how you can lower your energy consumption`
| 1.338 | `32155789598658560` | `Green lighting tips that will help you with your energy consumption! http://bit.ly/eyZvgc`
| 1.285 | `32890643206373376` | `ERCOT calls for energy consumption measures during Texas power emergency: http://j.mp/gYKYNk`
- `Detroit Auto Show`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.518 | `29249663391502336` | `speaking of freezing piss, is scamp at the Detroit Auto Show?`
| 3.147 | `29303621862559744` | `Detroit Auto Show: Mean and 'Green' http://ow.ly/3HOSe`
| 2.872 | `29220639558602752` | `Las promotoras del Auto Show de Detroit (fotos) http://dlvr.it/DwjYr`
| 2.659 | `30396111764066304` | `Final Numbers: 2011 Detroit Auto Show was bigger than 2010 in every way http://dlvr.it/F8rlN`
| 2.569 | `29177818524942336` | `It's sunny and a balmy 8 degrees. I-94 is clear no reason to miss the last day of the Detroit auto show.`
| 2.345 | `29525500988760064` | `#news #money Some cars just a dream for visitors to Detroit International Auto Show: Some cars... http://bit.ly/h1Nv5V #business #credit`
| 2.225 | `30356287451566080` | `The Green Hornet's car to make cameo at Detroit auto show: The car, known as Black Beauty, will be appearing at ... http://bit.ly/f7TnOP`
| 2.225 | `29327417629741056` | `Re: N.Y. Times on Detroit Auto Show: Chrysler says it is here to stay! [23/1/2011 19:59:34] - http://is.gd/WNKCOH`
| 2.045 | `29284877467652096` | `The auto show is freshh (:`
| 1.670 | `30224541766651904` | `Shenzhen International Auto Show: Shenzhen International Auto Show - Shenzhen International Auto Show Girls http://bit.ly/dSTxhM`
- `global warming and weather`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.586 | `33950240969068544` | `Weather feels great man I love "global warming"`
| 2.313 | `33142059846074368` | `All this snow and cold weather must be global warming. Yeah, that makes sense.`
| 2.313 | `33077737648300033` | `@imissTVXQ5 weird tlga ang weather.. Damn global warming.., hehe`
| 2.313 | `32878893287342080` | `Today's weather reminds me why I'm excited for global warming. #thankyoualgore`
| 2.119 | `30727128735547393` | `So...how's Al Gore going to spin the freezing weather? Global cooling due to global warming? What a con-artist.`
| 2.111 | `34644477629042688` | `RT @Boss_Wade: This weather rite here lets you know global warming is real lol`
| 2.028 | `34455776604983296` | `#YourTeamLoses & you blame the referees, the coach, the QB, the crowd cheering too loud, the weather, global warming, the illuminati...`
| 1.955 | `33990778749460480` | `GLOBAL WARMING UPDATE: Bizarre Weather, Destroyed Crops, And No More Right Whales http://bit.ly/dY3soP`
| 1.873 | `32861356432560128` | `GLOBAL WARMING`
| 1.717 | `33411350726180864` | `It is not Global Warming it is Global Jihad.`
- `Keith Olbermann new job`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.285 | `34829699771269121` | `Keith Olbermann's New Job Is... - Keith Olbermann knows how to get people talking - and tweeting! http://ow.ly/1behHV`
| 2.461 | `34989670014128128` | `What's Next For Keith Olbermann??? Announcing New Job Today ...: He's the Lindsay Lohan of pseudo- journalism. E... http://bit.ly/eD2YmZ`
| 1.953 | `35085871237709824` | `AP: Fired Keith Olbermann gets crappy job at some loser station that is seen by 20 liberals a night. LOL at loser Olbermann!`
| 1.942 | `34715311164887040` | `Keith Olbermann announcing new venture tomorrow http://huff.to/fJEpEC`
| 1.942 | `29725927323738112` | `#FOK News... Keith Olbermann's new cable news network!`
| 1.851 | `34944116676632576` | `http://bit.ly/dQ4YoD Keith Olbermann to unveil new career on Tuesday`
| 1.851 | `34800449370587136` | `RT @lgailey #Keith Olbermann announcing new venture tomorrow, Tuesday, 11 AM... // What?????`
| 1.851 | `29360350755295232` | `Hey HBO: Give Keith Olbermann a new show, STAT. Thanks much.`
| 1.801 | `34952094221860864` | `Current TV may become the new home of Keith Olbermann: Former host of MSNBC's "Countdown," Keith Olbermann, is s... http://bit.ly/g5xzCP`
| 1.801 | `34952093886320640` | `Current TV may become the new home of Keith Olbermann: Former host of MSNBC's "Countdown," Keith Olbermann, is s... http://bit.ly/efkx7W`
- `Special Olympics athletes`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.504 | `29528213755531265` | `Tar Heels host Special Olympics athletes: On Sunday, Special Olympics athletes got a chance to play on one of th... http://bit.ly/edtJgD`
| 3.199 | `29599881014280193` | `Tar Heels Host Special Olympics Athletes http://bit.ly/ft5kaj #UNCBball`
| 3.141 | `33445664922800129` | `Athletes compete, celebrate at Special Olympics Michigan State Winter Games: More than 1000 athletes from across... http://bit.ly/f1YSjD`
| 2.612 | `30526264926281728` | `Athletes prepare for Winter Games: The Winter Games is one of the five main Special Olympics competition events ... http://bit.ly/fWZHkg`
| 2.478 | `33396269518950400` | `Athletes compete, celebrate at Special Olympics Michigan State Winter Games: By Jordan Spence || February 03, 20... http://bit.ly/dVEHfW`
| 1.655 | `32657472703430657` | `POLL: What do you think the parking situation is at the special olympics?`
| 1.655 | `32145000863105025` | `Iron Butt Endurance Ride for Special Olympics`
| 1.641 | `30581469411803137` | `Special Olympics to provide needed boost http://bit.ly/hB5zMu #olympics`
| 1.599 | `31622888691867648` | `Teams splash in for Special Olympics: That's understandable because the Special Olympics of Louisiana, an organi... http://bit.ly/ghao02`
| 1.502 | `31029975851208704` | `@eonline Breaking News! Special Olympics http://theoriginalkirby.blogspot.com/2011/01/special-olympics_28.html?spref=tw …`
- `State of the Union  and jobs`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.493 | `29010511072862208` | `Sneak Peak to State of the Union: Jobs, Jobs, Jobs http://goo.gl/fb/XjGVg`
| 1.439 | `29536476286943232` | `State of the Union 2011: Jobs and More Jobs #president #stateoftheunion #jobs | http://dld.bz/HsHx`
| 1.333 | `29186381406277632` | `Reuters Top News : State of the Union speech to focus on jobs: Obama`
| 1.333 | `29035880375132161` | `Obama: State of Union to focus on jobs http://bit.ly/fvEx79`
| 1.271 | `29061881314410496` | `Obama: State of the Union speech to focus on jobs http://reut.rs/dSYzkl`
| 1.217 | `29878690959589376` | `The Economic State of the Union: America Is Already Creating Jobs http://bit.ly/eAuqmq`
| 1.217 | `29040795906482177` | `State of the Union speech to focus on jobs: Obama http://goo.gl/fb/iT5er`
| 1.169 | `28972912073510913` | `Obama: Jobs will be State of the Union's 'main topic' http://bit.ly/eQaVMw`
| 1.149 | `30150787321565184` | `The Early Show: State of the Union 2011: Jobs and More Jobs - White House Says President Will ... creating... http://tinyurl.com/4dt58g5`
| 1.127 | `30285454850924544` | `Jobs, Infrastructure Mark State of the Union http://bit.ly/g88n6o <-- #CRE, what are your thoughts? #REALESTATE`
- `Dog Whisperer Cesar Millan's techniques`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.543 | `30708594202648576` | `I am watching Dog Whisperer with Cesar Millan http://bit.ly/hRbSag (via @GetGlue) #DogWhispererWithCesarMillan`
| 2.543 | `29602995834454016` | `I am watching Dog Whisperer with Cesar Millan http://bit.ly/dzQ8xT (via @GetGlue) #DogWhispererWithCesarMillan`
| 0.808 | `30022191554764802` | `Dog Whisperer- Kenneling - http://ow.ly/1s0cZx #pets`
| 0.808 | `29373541111570432` | `Why do Filipinos call their kids like the dog whisperer calls dogs`
| 0.751 | `29255276477554689` | `DOG WHISPERER vs ITS ME OR THE DOG? | Cocker Spaniel Tips http://bit.ly/h5tJCZ`
| 0.731 | `30335800193323010` | `Dog Whisperer adds second show in Wilmington http://bit.ly/fZn6ku`
| 0.656 | `29688639214592002` | `Dog Training Group Class, Dog Whisperer Style www.K9-1.com | The ... http://bit.ly/dFmxSh`
| 0.571 | `30292051752914944` | `What am I the fog whisperer? I see another lost dog. This one is running around the construction area on La (cont) http://tl.gd/8dc7nf`
| 0.558 | `29195955517526016` | `There are still tickets available to see Cesar Millan in Red Bank NJ 2/25 8 pm HURRY!!`
| 0.509 | `29659593030242304` | `trouxe o livro do Cesar Millan pra ler nos horarios de folga... se é que vou ter! hauhauhua`
- `MSNBC Rachel Maddow`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.009 | `33550725875245056` | `Watching a Rachel Maddow video: The whole world is watching Egypt - http://on.msnbc.com/gnyGQr`
| 3.009 | `33461777790406657` | `لا يفوتكم التقرير "@ThamerSalman: Rachel Maddow rocks http://on.msnbc.com/9gUBdm #41417897"`
| 3.009 | `32824614375653376` | `Internet Hoax a flummox for MSNBC's Rachel Maddow http://sodahead.com/poll/1485945/`
| 2.899 | `33022351985614848` | `Watching a Rachel Maddow video: Protesters clash overnight in Tahrir Square - http://on.msnbc.com/e2Uw8r`
| 2.899 | `32988880349175808` | `MSNBC Rachel Maddow: Journalist reports they have WON Tahrir Square against the Pro Murbarak's THUGS! YIPPIE!`
| 2.516 | `33420944232034304` | `finally someone said the truth http://www.msnbc.msn.com/id/26315908/ns/msnbc_tv-rachel_maddow_show/#41417897 #Egypt #Tahrir #Mubarak #jan25`
| 2.489 | `32590294042025984` | `via @BluegrassPundit MSNBC’s Rachel Maddow Falls For Internet Spoof Story on Palin as Fact http://bit.ly/eese69 | #tcot`
| 2.426 | `33656631187210241` | `Highly recommended #podcast: MSNBC's Rachel Maddow Show. Video: http://ow.ly/3Q4jH Audio: http://ow.ly/3Q4jI #p2`
| 2.367 | `31825515270639616` | `RT @MMFlint: @maddow bcame the 1st anchor 2 say it! (She said it online;MSNBC had no news last night) So Wolf u share the prize w/ Rachel!`
| 2.357 | `32666646996844544` | `#nw The Rachel Maddow Show`
- `Sargent Shriver tributes`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.212 | `29089018561953792` | `Sargent Shriver Funeral Mass Held In Maryland (VIDEO): POTOMAC, MARYLAND (AP) -- R. Sargent Shriver was ho... http://tinyurl.com/4ryfywx`
| 1.835 | `28967095878287360` | `Mourners recall Sarge Shriver's charity, idealism (AP): AP - R. Sargent Shriver was always an optimist, pio... http://bit.ly/gqMcdG`
| 1.704 | `29208901270380544` | `USGOV Hundreds paid their respects Friday to R. Sargent Shriver, the first director of the Peace Corps, who died Tuesday. He was 95.`
| 1.653 | `29012766593384448` | `POTOMAC, Md. (AP) — R. Sargent Shriver was always an optimist, pioneering the Peace Corps and running the ... http://tinyurl.com/4f3y9fy`
| 1.564 | `29236353774391296` | `National Association of Community Health Centers: Remembering the Legacy of R. Sargent Shriver http://www.nachc.org/pressrelease-detail.cfm?PressReleaseID=642 …`
| 1.564 | `28968401699344384` | `Bono & Glen Hansard sang Make Me A Channel Of Your Peace at Sargent Shriver's funeral today http://wapo.st/geJQJc (via @StrongGirl)`
| 1.488 | `29355674039230464` | `U2's Bono on Sargent Shriver | National Catholic Reporter: Bono, the lead singer of the band U2 and a co-founder... http://bit.ly/gL6Tki`
| 0.599 | `29095373117063168` | `Heath Ledger fans light up internet with humble tributes http://bit.ly/hja4Qg`
| 0.555 | `33152944681517056` | `#IranElection Moving tributes inspired by events after #IranElection in Arts Showcase for #Iran http://bit.ly/frReEf`
| 0.548 | `29087268325036032` | `Top External Search Queries for Jan 22 2011: forrest sargent`
- `Moscow airport bombing`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 4.717 | `29565853112467456` | `Moscow airport bombing :S`
| 3.729 | `29578335377952768` | `Bombing tragedy in Moscow at the airport: http://bbc.in/fYM4Ai`
| 3.729 | `29573016115478528` | `Bombing reported in Moscow DME airport; prayers to all the families.`
| 3.516 | `29622471892140032` | `May all the victims of terrorist bombing at #Moscow airport Rest in Peace!`
| 3.516 | `29585829366071297` | `Prayers and thoughts go out to the victims of the Moscow airport bombing. Tragic.`
| 3.442 | `29550342844715008` | `Suicide bombing at Moscow airport http://www.wordtravels.com/TravelNews#moscow_bomb`
| 3.336 | `29659332027088897` | `Map and diagram of the airport bombing in Moscow. http://nyti.ms/hX9d5I`
| 3.180 | `29744040551391232` | `Bombing at Moscow airport called terrorist attack http://bit.ly/fHDRaM`
| 3.180 | `29590400175964160` | `Terrorist bombing @ Moscow airport. This makes me so sad. I'm speechless. Moment of silence.`
| 3.114 | `29598054659133442` | `Video of Moscow bombing aftermath. Sad -- First witness video moments after Moscow Domodedovo airport bombing http://www.youtube.com/watch?v=BUQPlEr1puI …`
- `Giffords' recovery`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.047 | `31020191437557760` | `Watch: Road to Recovery: Gabrielle Giffords' Green Light? #all`
| 2.956 | `30527259790020608` | `Doctors: Giffords' recovery 'not a sprint': In the days and weeks to come, Rep. Gabrielle Giffords' recovery fro... http://bit.ly/dRhUPt`
| 2.873 | `30658867457167361` | `CNN: Doctors: Giffords faces long, rocky road to recovery`
| 2.873 | `30633783929606144` | `Doctors: Giffords faces rocky road to recovery: In the days and weeks to come, Rep. Gabrielle Giffords' recovery... http://bit.ly/i9gHxT`
| 2.797 | `30604291001556992` | `Doctors: Giffords faces rocky road to recovery - In the days and weeks to come, Rep. Gabrielle Giffords' recovery fr... http://ow.ly/1b3hiO`
| 2.599 | `29609129211330560` | `Giffords' Recovery Continues Despite Complication http://goo.gl/fb/crNf3`
| 2.599 | `29597334014787585` | `Giffords: Rocky Road to Recovery http://j.mp/hTjn6a #health #news`
| 2.599 | `29524599578304512` | `Brain Injury Recovery That Lies Ahead for Giffords - http://nyti.ms/hLJTBc`
| 2.599 | `29018577633804288` | `Rep. Sullivan Talks Giffords' Road To Recovery http://bit.ly/fLDPeo`
| 2.550 | `30631645262385152` | `iphone Doctors: Giffords' recovery 'not a sprint': In the days and weeks to come, Rep. Gabrielle Giffords' recov... http://bit.ly/13rtYI`
- `protests in Jordan`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.762 | `31396975362445313` | `After #Egypt - Yemen, Jordan on the brink of #Uprising. Many protests underway.`
| 1.612 | `30701998034653184` | `EoZ: Protests continue in Egypt; new protests in Yemen; Jordan next? http://goo.gl/fb/SaRhS`
| 1.593 | `31772901715746816` | `Pro-Democracy Protests break out in Jordan - http://tinyurl.com/4bouwqh - catspirit`
| 1.580 | `29367743086075904` | `PAK Protests erupt in Yemen, Jordan, Algeria and Albania against Govt.: Activists in Yemen, Jordan, Algeria and ... http://bit.ly/fqImyT`
| 1.412 | `32160881890562048` | `RT @nationaljournal: Protests beyond #Egypt: Photos from Yemen, Lebanon, Jordan and more. http://ow.ly/3Nv8o`
| 1.365 | `32142111931375616` | `NEWS: #Egypt Demonstrations Inspire Protests to Erupt in #Jordan, #Yemen and #Sudan. http://ow.ly/3NsXW #jan25`
| 1.365 | `32141682078126080` | `NEWS: #Egypt Demonstrations Inspire Protests to Erupt in #Jordan, #Yemen and #Sudan. http://ow.ly/3NsXW #jan25`
| 1.365 | `31765380544331777` | `VIDEO: Democracy protests in Amman, Jordan on Friday, Jan 28 http://post.ly/1YLtN  #jan25 #ReformJO`
| 1.321 | `31068649536102400` | `You bet. RT @CrowleyTIME: Modest protests in Jordan, too. http://bit.ly/eAg2FB Must be uncomfortable in Israel right now.`
| 1.321 | `31055600443981824` | `Protests and uprisings in Egypt, Jordan, Tunisia and Albania-- pretty much surrounding us. More importantly, I can't find my chapstick!`
- `Egyptian curfew`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.833 | `31045442808586240` | `Egyptian president extends curfew to cover entire country`
| 2.671 | `31010567162433537` | `BREAKING: Egyptian TV announces immediate curfew in Cairo #Jan25`
| 2.651 | `31044476419964928` | `#Mubarak forces curfew over ALL Egyptian cities, and still no appearance to give the announced speech. #Egypt #Jan25 #Protest #Curfew`
| 2.567 | `31367822701494272` | `Egyptian protesters defy curfew: Tens of thousands of demonstrators defy a curfew and remain on the streets, des... http://bbc.in/f3ZCNH`
| 2.534 | `31046710490832896` | `RT @BBCNews: Egyptian president extends curfew to cover entire country`
| 2.420 | `31096457457041408` | `Egypt`s Mubarak declares nationwide curfew: Egyptian President Hosni Mubarak on Friday extended a curfew to cove... http://bit.ly/gVUZtI`
| 2.416 | `32169781150883840` | `LMAO! “@hmikail: The only Egyptian citizen obeying the Curfew is Husni Mubarak. #Jan25 #Egypt”`
| 2.416 | `30010192120782849` | `The Egyptian authorities may impose curfew in the coming hours according to Al Jazeera`
| 2.313 | `31044465690939392` | `Egyptian military deploys in Cairo under curfew http://usat.ly/gteGBy  via @USATODAY`
| 2.223 | `31725521607856128` | `@FreeAdviceMan1 If I were an Egyptian in Egypt I would be breaking curfew & in the streets demanding my rights too. It's the people's time.`
- `Beck attacks Piven`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.546 | `32174687102435328` | `Glenn Beck's virulent attacks on CUNY prof Fox Piven are inciting threats; AAUP says STOP the irresponsible rhetoric: http://bit.ly/dG2FrV`
| 1.692 | `30062538221686784` | `Glenn Beck target Frances Fox Piven: Beck is 'very scary' http://huff.to/fJ0oDR via @huffingtonpost`
| 1.580 | `30161016113340417` | `Amazing, Giffords gets shot but Beck's the real victim. Piven gets death threats (on Beck's site!), but Beck's the victim. #pityparty`
| 1.489 | `29889791659081728` | `RT @Lady_grrrr: Beck, Piven & the Way the Left Works - http://bit.ly/hX5BLw`
| 1.461 | `31569527422586881` | `Frances Fox Piven defies death threats after taunts by anchorman Glenn Beck http://www.guardian.co.uk/media/2011/jan/30/frances-fox-piven-glenn-beck … via @BrkingWorldNews`
| 1.435 | `29669058383183872` | `Of course... the NYT article about Beck/Piven is part of the Soros conspiracy to destroy America, or something. #GlennBeck #Stupid`
| 1.435 | `29020656456699904` | `New post: Glen Beck and Frances Fox Piven on violence http://www.teapartynews.net/45`
| 1.401 | `30865819902672896` | `#MSNBC Host Who Called Beck a Nazi Attacks Beck for Nazi Analogies http://www.theblaze.com/stories/omg-msnbc-host-who-called-beck-a-nazi-attacks-beck-for-nazi-analogies/ … #cnn #chicago #p2 #tcot #tlot #twisters #sgp #ocra`
| 1.386 | `30023535518818304` | `Glenn Beck isn't going to silence Frances Fox Piven that easily. http://bit.ly/gPswMt`
| 1.386 | `29584060128960512` | `Glenn Beck has posted a story on Frances Fox Piven at his flaming website. He just incites and incites and incites. What a lowlife!`
- `Obama birth certificate`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.240 | `29267269871734784` | `Obama Birth Certificate Cover-Up Continues http://bit.ly/htjPP8`
| 2.966 | `30573952518586368` | `Hot Update: Abercrombie Birth Certificate: Answering the Eternal Mystery of Obama's Birth http://bit.ly/ibwZzP`
| 2.957 | `32947526336249856` | `Beck is right, we need to look in Cairo museum for Obama's birth certificate #p2`
| 2.957 | `30781723516076032` | `Abercrombie Admits Failure To Discover Obama Birth Certificate http://su.pr/29MAjE`
| 2.957 | `30343979849482240` | `Obama birth certificate 'egg on face' for guv http://goo.gl/fb/1KWvd`
| 2.957 | `29996599845322752` | `Limbaugh Demands Obama Release Birth Certificate http://onlywire.com/r/21885898`
| 2.957 | `29996506652082176` | `Limbaugh Demands Obama Release Birth Certificate http://onlywire.com/r/21885904`
| 2.957 | `29881885689643008` | `Hawaii official now swears: No Obama birth certificate. http://bit.ly/fNpBU1`
| 2.957 | `29857973861879808` | `Obama story about Hawaii past and lost birth certificate: http://bit.ly/gRA9y4`
| 2.855 | `30035763173265408` | `There is no Barack Obama birth certificate in Hawaii, absolutely no proof at all that Obama was born in Hawaii. http://bit.ly/hJ3lTW`
- `Holland Iran envoy recall`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 0.903 | `34714824982134784` | `Dutch recall envoy in Iranian row - http://www.bbc.co.uk/news/world-europe-12381284`
| 0.876 | `34546003663126528` | `Breaking: Dutch recall ambassador from #Iran over execution of Dutch-Iranian Zahra Bahrami, summon Iran ambassador`
| 0.844 | `29517561792045057` | `Iran ready to comply with Tehran Declaration terms - Iranian envoy to IAEA http://bit.ly/dJmjlu`
| 0.789 | `34301010126049281` | `Envoy Calls for Iran-Venezuela Joint Confrontation against Imperialism - Fars News Agency http://bit.ly/gBMqlx #news`
| 0.494 | `34171092457234432` | `Goodmorning Holland`
| 0.494 | `33142912841678848` | `@tommcrae And Holland?`
| 0.494 | `32473946146213890` | `The Holland marsh`
| 0.494 | `30591208233373696` | `@holland bakery`
| 0.430 | `32920337112891392` | `Iran?`
| 0.428 | `31658422512386048` | `@AshleighJackson No Holland why>`
- `Kucinich olive pit lawsuit`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.343 | `30730275864449024` | `Enough about alien Kucinich and the olive pit. That is all.`
| 3.332 | `31468826088439809` | `Kucinich Settles Lawsuit Over Olive Pit http://bit.ly/fay1F6 That was quick... I guess he couldn't stand himself any longer #tcot`
| 3.238 | `31120562348625920` | `BREAKING: Dennis Kucinich settled his lawsuit ~Oh, thank god. The olive pit is off the tooth—I mean, hook... http://slate.me/iaYMxt #usguys`
| 3.072 | `30349976752099328` | `Rep. Dennis Kucinich files lawsuit over olive pit - ... seeking $150,000 in damages - USATODAY.com: http://usat.ly/eC17hV via @addthis`
| 2.469 | `31261786745339904` | `Ohio: Kucinich Settles Suit Over Olive Pit http://bit.ly/ifDhBr`
| 2.469 | `30824505316220928` | `Dennis Kucinich Sues over an Olive Pit in His Sandwich http://ht.ly/3LHlk`
| 2.364 | `30366659071975424` | `Finally!! A trending topic worth trending! Dennis Kucinich and the Perpetually Pesky Olive Pit. :/`
| 2.364 | `30364026441572352` | `Dennis Kucinich Sues Congressional Cafeteria Over Olive Pit http://gaw.kr/eiuDIH`
| 2.271 | `30410427095588864` | `Rep. Kucinich sues over olive pit in sandwich (The Arizona Republic) http://feedzil.la/hcA1vB`
| 2.271 | `30347151305342976` | `Kucinich sues Longworth cafeteria over olive pit in sandwich: report http://bit.ly/ig3jCg`
- `White House spokesman replaced`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.657 | `32871752056971264` | `White House spokesman Robert Gibbs refuses to call Mubarak a dictator when asked by reporter #egypt #jan25`
| 1.550 | `30770417580904448` | `AP source: Obama picks Biden's spokesman for White House press secretary http://bit.ly/eqTX60`
| 1.422 | `29695515914407937` | `#WeJew_Jewish White House Spokesman Reads President's Comments on Moscow Terror Attack: - http://bit.ly/dWoP1K #Video Sharing`
| 1.386 | `30755695221547009` | `VP Joe Biden's spokesman Jay Carney named White House press secretary as staff chief Bill Daley unveils personnel moves.`
| 1.386 | `30751625622585344` | `Politics: Former reporter Carney next White House spokesman \n (Reuters)\n: Reuters - Jay Carney, a communica... http://bit.ly/gTNAS4`
| 1.186 | `30094484930826240` | `white house white house`
| 1.148 | `29361798910050304` | `white house white house white house ..#swag`
| 0.889 | `29706615770845185` | `The clothing house formerly known as "the floor" has been replaced with "the dresser." Go team. No more bachelor living. Well for a week.`
| 0.839 | `29626070688866306` | `#becauseimblack I should rename the White House.`
| 0.839 | `29578134802145280` | `#becauseimblack I should rename the White House.`
- `political campaigns and social media`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.814 | `31471973502558208` | `Social Media Campaigns http://bit.ly/fDDA87`
| 1.547 | `29921958300753920` | `Social media campaigns 'require clear goals' http://bit.ly/eyR1T4`
| 1.524 | `29836714683924480` | `The Political Power of Social Media http://is.gd/M4cofF`
| 1.481 | `30440776995446784` | `I've been really busy building social media marketing campaigns for my friends at @rfamilyfarmil and @kokopellispub.`
| 1.423 | `30752353867014144` | `Musings about librarianship: 4 Successful social media campaigns for and by libraries http://ow.ly/3LxjI`
| 1.423 | `29769447698857984` | `I am slamming these social media campaigns out! Tweet, yes! But make sure your business is balancing it out with a campaign & interaction :)`
| 1.423 | `29233816241446912` | `Musings about librarianship: 4 Successful social media campaigns for and by libraries http://bit.ly/eQ2l5T`
| 1.325 | `32274142413717504` | `Super Bowl TV ads to unleash social media campaigns - Lost Remote http://bit.ly/guq6u8`
| 1.325 | `32274140148797440` | `Super Bowl TV ads to unleash social media campaigns - Lost Remote http://bit.ly/hGMVHJ`
| 1.325 | `29424217959174144` | `This just in Bees Awards: World's best social media campaigns in the spotlight http://ow.ly/1b01Am`
- `Bottega Veneta`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 3.945 | `33207140613103617` | `Bottega Veneta to launch first fragrance! - http://newzfor.me/?c1dl`
| 3.743 | `33132476125675520` | `Beyonce On The Go… with hers "BOTTEGA VENETA BAG" http://nblo.gs/dPbFh`
| 3.653 | `33981481269338112` | `Donne Rumene: Bottega Veneta Collezione Scarpe Primavera Estate ... http://donnerumene.blogspot.com/2011/01/bottega-veneta-collezione-scarpe.html?spref=tw …`
| 3.569 | `30604172848009216` | `Leon payn for bottega Veneta coming soon. http://ow.ly/i/7A93`
| 3.569 | `30316236273360897` | `RT @vogue_italia: Bottega Veneta against unemployment http://ow.ly/1b25qH`
| 3.417 | `33967238587359232` | `Bottega Veneta bookbag and Louis Vuitton purse... It doesnt get any better than that... I am so grateful for my #blessings`
| 3.163 | `33811002248138752` | `Elle Fanning, Rodarte Muse... Bottega Veneta's Palm Beach Bag... - http://newzfor.me/?cj5n`
| 3.163 | `33464465403940864` | `Elle Fanning, Rodarte Muse... Bottega Veneta's Palm Beach Bag... - http://newzfor.me/?cj5n`
| 3.163 | `33463500768542720` | `Elle Fanning, Rodarte Muse... Bottega Veneta's Palm Beach Bag... - http://newzfor.me/?cj5n`
| 3.163 | `31950180282540034` | `Bottega Veneta teal and turquoise knot intrecciato snakeskin clutch http://f.ast.ly/CXWrG`
- `organic farming requirements`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.538 | `32273499913453568` | `Farming : New Congress shows hostility to organic farming | Rodale Institute http://bit.ly/hPLF7l`
| 1.481 | `33000974503120896` | `Ramsey Bros have a wide range of industry leading equipment to suit your farming requirements with branch... http://fb.me/GE8rvk0u`
| 1.478 | `33600100227883010` | `Organic Farming Changes Everything for a Community in India | Gaiam Life http://life.gaiam.com/article/organic-farming-changes-everything-community-india …`
| 1.444 | `33964088472117248` | `USDA/Vilsack decision on GM alfalfa WILL HURT ORGANIC FARMING in this country!`
| 1.444 | `32496916352737282` | `Is organic farming policy-driven or consumer-led? http://bit.ly/gs3v7f`
| 1.444 | `29014865242759168` | `Interested in Organic Farming? Then do one of these cou... - http://fwix.com/a/96_4796c09938`
| 1.328 | `32878274442956800` | `RT @nontypicalmom2011RT @troptraditions: We are committed to family farming & organic standards. We stand firmly against #GMOs. Food ...`
| 1.286 | `32117741028638720` | `Organic eggs set for testing times - Farming UK: http://dld.bz/J3ru - we must try and help organic producers wherever possible`
| 1.237 | `29497356420390912` | `Vt. organic farming group prepares for conference http://www.wggb.com/Global/story.asp?S=13894269`
| 1.197 | `30756753977118720` | `@kellyjanice Organic farms cannot produce enough food to feed the country. Monsanto crops are vital to the (real) farming that feeds the US.`
- `Egyptian evacuation`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.804 | `29646008031911936` | `Fire & Evacuation Drilll .`
| 1.473 | `29394389834932225` | `The whole Katrina evacuation plan! #FAIL`
| 1.275 | `29837667814342656` | `Updates !: The Evacuation - Sacrifice Nothing - SC2 Achie... http://starcraft2.igoofer.com/the-evacuation-sacrifice-nothing-sc2-achievement`
| 1.088 | `31423721797324800` | `Gas leak at Texas oil refinery prompts evacuation http://dlvr.it/FMXmN`
| 1.088 | `29626932115013632` | `Fires, gas leaks prompt evacuation of Fairport Harbor http://lnkd.in/5XVfJG`
| 1.041 | `29533280697065472` | `Long lines of traffic leaving Fairport Harbor. Full village evacuation underway @WEWS`
| 1.000 | `31707575615496192` | `We have a multi-national team, positioned in Cyprus, that can effect evacuation plans for private clients in #Egypt.`
| 1.000 | `29935288641921025` | `#WWII 25 January 1945: German evacuation of military personal and civilians from East Prussia begins.`
| 0.964 | `29400901689024512` | `People left their blankets and chairs on the beach for the evacuation? It wasn't like the water was receding already. #H50`
| 0.931 | `31896461222346753` | `Minister #Cannon on the situation in #Egypt and evacuation plans in place for #Canadian citizens: http://ow.ly/3MY7G #canada`
- `carbon monoxide law`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 2.746 | `32005451423948800` | `Law requiring carbon monoxide detectors goes into effect Feb. 1: Beginning Tuesday, a new state law takes effect... http://bit.ly/hnxJp7`
| 2.517 | `30783835524300800` | `just your average carbon monoxide leak......`
| 2.055 | `32307745206042624` | `I am afraid of carbon monoxide. (Gee. I don't know where that came from!)`
| 2.055 | `29725732527669249` | `Dying of carbon monoxide poisoning. http://yfrog.com/h8faqjsj`
| 2.055 | `29653370469879808` | `Carbon monoxide warning to homeowners http://ow.ly/3JoxH`
| 1.859 | `32569981321347074` | `EPA proposes no changes to carbon monoxide limits. http://nyti.ms/i1QtNB`
| 1.859 | `31181287960084482` | `Family Suffers Carbon Monoxide Poisoning From Running Car http://tf.to/fzdU`
| 1.859 | `30840353938477056` | `Carbon Monoxide Sickens Two Kids In The Bronx http://bit.ly/g62H4C`
| 1.859 | `30721013201240065` | `@BT Wow what a story!!!!!! ..I'm so happy you all are ok ...carbon monoxide is a silent killer ...`
| 1.710 | `30053106599329793` | `#Tech #TechNews Carbon Monoxide Could Have Blown Up Comet; Could Repeat http://bit.ly/fuz9HH #DhilipSiva`
- `war prisoners, Hatch Act`

| Score | Tweet ID            | Tweet content
|-------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------
| 1.153 | `31939767742496769` | `Political prisoners in Iran - act now - http://bit.ly/7WrH87 - #iranelection #iran #fb`
| 1.128 | `29751056560820224` | `Hatch Act protects Civil Servants from abuse/use for political campaigns. Schdle C's are political appointees -`
| 1.128 | `29729125346312194` | `huffingtonpost Bush White House Violations: Political Briefings Reportedly Violated Hatch Act http://huff.to/gquujn`
| 1.002 | `29723425576587264` | `Report by special counsel's office finds Bush White House engaged in Hatch Act violations - Minneapolis Star Tribune: WASHINGTON - Th...`
| 0.953 | `30170530963259392` | `Hatch Act probe: Bush broke law to help Fitzpatrick: The White House illegally used federal funds to s... http://bit.ly/e1kLnX #politics`
| 0.405 | `30506153502834690` | `War again`
| 0.391 | `33219512429977600` | `@FeliipeNasc ook, so act`
| 0.362 | `30547566101798912` | `The great war is a spiritual war.`
| 0.347 | `34508307162992640` | `#iranelection Information & action on condemned prisoners in Iran http://j.mp/9HqzMG  #Iran #HumanRights /@DokhtarGol`
| 0.340 | `32498873666637825` | `‎"Senator Hatch was the first person in the Senate to call the individual mandates unconstitutional."-Carl Cameron http://dld.bz/JE9A #utpol`

### Include a sample of 100 tokens from your vocabulary.

* `phonepag`
* `series2`
* `36fo`
* `hdii16`
* `url`
* `huahaha`
* `joolie18`
* `flamez`
* `dorasbeag`
* `gvcak8`
* `bwahahahaha`
* `hln`
* `pwgdochi`
* `action-joe-the-story-of-the-french-gi-jo`
* `whitstabl`
* `kappalumaackel`
* `viki`
* `gqndpe`
* `virgodiv8`
* `2nvep`
* `dea`
* `gjikuh`
* `suppor`
* `frankieedgar`
* `ihb33t`
* `enysa3`
* `braba`
* `1b8yci`
* `gantengan`
* `manoww`
* `streamlin`
* `miskeenfresh`
* `e4goq`
* `nowaday`
* `fw8a03`
* `suportg`
* `f3zrbs`
* `1600`
* `bobsled`
* `f2kjff`
* `croix`
* `internazional`
* `bonjour`
* `righttt`
* `saludo`
* `weed-fir`
* `cz7c`
* `descriptionfrom`
* `immort`
* `twitvid`
* `ar2011012204023`
* `h8ptp1`
* `şifrelem`
* `p3k`
* `lindsaylohan`
* `люди`
* `under`
* `forev`
* `hedeu9`
* `timminchin`
* `hpjikd`
* `πως`
* `connector`
* `manjud`
* `6d6pr`
* `11-01-28`
* `graaaaaaaaaay`
* `fijdlf`
* `fsivm8`
* `progr`
* `bandung`
* `wareh`
* `dylantwm`
* `prioritize-your-website-not-your-social-media-account`
* `business-12389416`
* `f7gli9`
* `15695034`
* `joan`
* `369`
* `hotair`
* `cumbrian`
* `قسم`
* `juist`
* `5jib`
* `mpfapp`
* `1b1cuf`
* `jacqu`
* `fug`
* `radio4`
* `fidarab`
* `hcctkt`
* `watcin`
* `urchin`
* `gesehen`
* `cogener`
* `ame`
* `artista`
* `paker`
* `creator-destructor`
* `2987oc5`


