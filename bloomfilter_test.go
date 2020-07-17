package bloomfilter_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/R0L/bloom-filter"
	impl2 "github.com/R0L/bloom-filter/calculate/impl"
	"github.com/R0L/bloom-filter/storage/impl"
)

var bf *bloomfilter.BloomFilter

func init() {
	var length uint32 = 1000
	bf = bloomfilter.NewBloomFilter(impl.NewFileStorage(length), impl2.NewFnvCalculate(), impl2.NewCrc32Calculate())
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestBloomFilter_Add(t *testing.T) {
	var hStrs = []string{"HNd)~U*@$IjkctDp8k3HZH^hTiM#vG&S@zV)rZmS9@8n*NZN", "1wUiFsILXx)11jZ4qYSu", "jkM8nVR*eyanBJBm*noFEcN72#%", "Tb4QkZSMn2eiYktA$G$4JL&LWUYctxIIKDmvOQZRO%", "L0NN(OiC*!d1KJjd6%v57&hp6C", "Va836XA7PuT)~!_QlmyOWt1O5TXv~K^2U1", "B((eDScT)!b@P~Eugu*fXRrLevK", "mz(dfmYC4DwpV%+Xmuvn3tK", "@gxQ)kFHm!a2LBl1GF@TA", "TNaNh#H36xD1I", "e6h$_t7QT7Q9$(j9", "ohtMV*wArBYHp~T*bGlp_qBih)+tEjRBaZc*23Yv*b1oM*~T", "Ivf!jf@#6I0rdgpu)wWdt06nOqBMh^v5xp", "p#VYpeomNe+fsY0j", "~OsBHyvuG^_c46j$%#aLe5)HcjlRSNAnyv9e7LZ4px@EF4k7T0", "T)#87d77B3DqD^DCeD)^1K_", "$Fw0icm()^3a^BU6B(1a3Qh", "xLWaSI&huf!_gmy@h", "$w%vN8T87$C@%put@+SgKHieqg9gh", "OMnojUh**x8!iKIEn^*oVg", "KKYQ%hs7uif4zRL_%E6T*LV^wKGXaz45+S", "OeFT~F8%4hSueD82Sj&Da8hd9t9t1vvMtiYIYd8zlk^", "n^EAlKF&n(ITx+@9eDCQf6", "GyCTCTpxF0PtWaaPM5w2W7iPfMYTYF5", "ZklU_#fls^mvyVQMFKkR", "7@_BLb3AaBl@TJWo#3Df23AJTb2u", "lqN!AOtih_y#m3(@+u~#@x8ONs2Bj", "GCKk2*y70vW$", "fRE^Y6__13tE6G", "QkpM6AKB8ku6yJISGe~R5TcNqddFw~td_24", "MQHmQp5@v)mLPcvt(w)eV_fjn*d468H1kV4BQ78", "!elJLIiU_RnFXecKYn^_d!3eGz98h3G0be%#BCW", "^assNEhG(X0FUM)M8SZKyCb2c8V8B+~~o0jQPUNT$", "OsaKp20tjD!qy#bU", "Ds(DfH9$Q7zwr_E70xLuVUjGnbqlZeK4NgUB+BdwwFPJE", "niXl3s)aL4hjuvfzbzteXrnqu8At*9gEwA5gjfIa8~", "WAeTdDj554jZeiR1uYi0rMMW!Py$OZjEi4bYg^1n", "uiA44S$Uoe#v", "B5qaSBOE2%en!q&BA)&MY", "85YOLufpF_d1GkpeC", "+3sWCsWKwf!xThg7_TsqaP2N@Bcr37rD~5Tz%*", "~Mo&tjeF%OX)B_Tkx", "e59er)7h37zBeZ9p", ")W0Wy2XjR9t1", "GzUALW*F1toLMP+wCcY5#XKQRB)hK1", "exW7s7yIAIMYcIBJIr~Z!xpM5)vpAB0+b", "0fEF7sur^_lb!1a@xw$*w^*02!WTHf5NL%!szl4N2C@^)*(@", "~49H5BgWOGHrUte(vnOvkkjk#j$C_", "K5v2Rt#T~n+BTGDej^5RHiazwTm", "LZ5EsZKukv#wD_d&WaZ5y^))", "+KM$PPk4M(R%6I)K^~Rl#s%wds3PjuWw&@^NQPCseD~uZ", "v0Jgtg&n)vyN9qLy8Rjrn39LHZo%HnDAyzKuryPiqDGk~Lgx", "+WaCUBHEx+eIU~_kI@$WQ%", "xc_xjNVBDE~*k%^wkC^~1Tq7", "VyTe(4f#0NBl5kNL1Jag33Z9LbZ45Mx95Jq!", "_OEYUEx$N8Q1RWN~Sj~$0)A~(T^9TZ~H*S~+qL", "+jTlG9t!9D", "p4KY+7(nCvtRd%)blvwkn", "%Tn22i)Ld(ne#OcVYMb+j~4q6z%ef_!L$t#PBm8eU!5", "q5ck43n8l0oeW^Rt2q58(z~muE", "GV&!i+ndHlY_1NAJ#k9BWeVBYgB7H8gyJoV*LC)fs#e8", "+tL9N_(7FXFMkN3xNt~aIvZ_AdE~w1&e", "x(WADb2H)6#f8@c9)wV~5F!ZIMdI*gM2", "uO%akaDfUwj*(HjnC5(TuMYjhTzlw!OK4r", "FMrR1G+Nb&cPx(a3y)d_dZukD4_SvAKcTC(VGR1T", "I0YJLXvd7hN&1Vnaef_jVP^i~C*AUq", "50@~EUpkSXHNIbJ9x!dJ#~OD%6#(XE&_S", "YNw3qmyPP8@6G5&Sa!_Ah29hc!!HHg9!^QkV_O8)n~1", "(G+i^c^(1!FjSG+^%8D6qs4+*fkr^1O6vO", "m(R&RP5XGHd27@8E6", "c!7d$Eby90b2TR4Q%nF9LIXDOXtA+%kR%g7k*", "BiR*qZwotviSljc!L2HLpW23XJTrH%M@D~wZLVKlg7o&TP%M", "~HQwfmmsekrdZ$b", "m6+JSU2sZMwb!d4a", "7+GAfqYYH&rM(QtEizta_iw", "@OP(hLYUFQb$S6M_MH2hcv+sAwEvNb~V6VK*AO_AA9wE", "bW+AWVXFn6IwThuGAFyujcdXouKtM#v!rTaxZc$UTwD&yNc#P", "cbwVB@k@@sl^", "T~XRdgVO90me~6MWRr~W$Wh", "lLbY)uTf1*G2C*Xb28inpC58^vvz4OzM(FRdr)%k!sBAf3M(", "u*H6e0)0!v2g7(k6#1HLQ6gkDm", "a7%q9&Dj97", "!(P3kjzkzG+nG5kGPsxwpIN7#lM8wK&R)xotP+bpP#m2TdjIZf", "^L9s@MZDdWZxfZqgs_6EtFAXTaCHSn~n", "8p!u_%suwQ0nf!NDzQPpCpk_*t6VgYck(a0(hI~kPbnge_hY", "aYvMCCkI4OIPJDLRVHVe2oviDh!(IpFnK", "7svk)Y%tVG5hRt%3yg8Pg@gEvXoN2*Xw33wNf5wK*QpRJnb", "mh3nnraEMpFDPFi_a~JY8b#xOIUFAAxU!j^Pj", "r0j5bPyN_#RR5It%3Tlqb9+w*nvoU", "TdI8V8t8YTzOr0^))", "gS1%wx^ii~obaFtfTJjqBAbnw25_rivwINpZq2kS%tmy(iGQYo", "+RF*ctkiT%FUP6!66B08(%B", "QSS0gnHuqw^ioBT35ZsYopBn*wrgvAEMI42&&op#BB^Yz", "D2qO#zQmWc1DBV2$g)81r5Y^4ZRxikJ0dDlBG+av", "eX^D^26(wIt*M!i*3QlMIpRxu$Y@", "o%_ACwx6(p3Mf@7U", "Ixmr@Evx(7N7X!s&x(", "y0BBZ6)^Vx(Ao%ea0B3", "M8m(SrDNatIz5ce83g~vNwpcE@aH#yk_d^W4X4*OHis6OX", "LCtC%wX@TzcWWms0S&e"}
	for _, hashStr := range hStrs {
		err := bf.Add([]byte(hashStr))
		assert.Nil(t, err)
	}
}

func TestBloomFilter_Check(t *testing.T) {
	TestBloomFilter_Add(t)
	var hashStr = "HNd)~U*@$IjkctDp8k3HZH^hTiM#vG&S@zV)rZmS9@8n*NZN"
	err := bf.Check([]byte(hashStr))
	assert.Nil(t, err)
}

func TestNewBloomFilter_probability(t *testing.T) {
}
