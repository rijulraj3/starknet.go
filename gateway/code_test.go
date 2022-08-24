package gateway

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/dontpanicdao/caigo/types"
	"github.com/google/go-cmp/cmp"
)

const codeDump = `{"bytecode": ["0x40780017fff7fff", "0x1", "0x208b7fff7fff7ffe", "0x400380007ffb7ffc", "0x400380017ffb7ffd", "0x482680017ffb8000", "0x3", "0x480280027ffb8000", "0x208b7fff7fff7ffe", "0x20780017fff7ffd", "0x3", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480080007fff8000", "0x400080007ffd7fff", "0x482480017ffd8001", "0x1", "0x482480017ffd8001", "0x1", "0xa0680017fff7ffe", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffffb", "0x402a7ffc7ffd7fff", "0x208b7fff7fff7ffe", "0x20780017fff7ffd", "0x4", "0x400780017fff7ffd", "0x1", "0x208b7fff7fff7ffe", "0x400380007ffc7ffd", "0x482680017ffc8000", "0x1", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x48297ffc80007ffd", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffffb", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x43616c6c436f6e7472616374", "0x400280007ff97fff", "0x400380017ff97ffa", "0x400380027ff97ffb", "0x400380037ff97ffc", "0x400380047ff97ffd", "0x482680017ff98000", "0x7", "0x480280057ff98000", "0x480280067ff98000", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x47657443616c6c657241646472657373", "0x400280007ffd7fff", "0x482680017ffd8000", "0x2", "0x480280017ffd8000", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x476574436f6e747261637441646472657373", "0x400280007ffd7fff", "0x482680017ffd8000", "0x2", "0x480280017ffd8000", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x47657454785369676e6174757265", "0x400280007ffd7fff", "0x482680017ffd8000", "0x3", "0x480280017ffd8000", "0x480280027ffd8000", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x53746f7261676552656164", "0x400280007ffc7fff", "0x400380017ffc7ffd", "0x482680017ffc8000", "0x3", "0x480280027ffc8000", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x53746f726167655772697465", "0x400280007ffb7fff", "0x400380017ffb7ffc", "0x400380027ffb7ffd", "0x482680017ffb8000", "0x3", "0x208b7fff7fff7ffe", "0x400380017ff97ffa", "0x400380007ff97ffb", "0x482680017ff98000", "0x2", "0x208b7fff7fff7ffe", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x2", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffffe", "0x400780017fff8000", "0x0", "0x400780017fff8001", "0x0", "0x48127ffe7fff8000", "0x208b7fff7fff7ffe", "0x20780017fff7ffc", "0x5", "0x480a7ffa7fff8000", "0x480a7ffd7fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x482680017ffc8000", "0x800000000000011000000000000000000000000000000000000000000000000", "0x40337fff7ffb8000", "0x480a7ffb7fff8000", "0x480a7ffa7fff8000", "0x480a7ffd7fff8000", "0x48317ffd80008000", "0x400080007ffd7ffe", "0x480080007ffc8000", "0x400080017ffc7fff", "0x482480017ffb8000", "0x1", "0x482480017ffb8000", "0x3", "0x480080027ffa8000", "0x20680017fff7ffb", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff8", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x2", "0x480a7ffa7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x480280007ffb8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffe2", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffd5", "0x40137ffd7fff8000", "0x480280017ffb8000", "0x40297ffd7fff8001", "0x48127ffb7fff8000", "0x48127ffc7fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x2", "0x480a7ffb7fff8000", "0x480280007ffc8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff6e", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffc6", "0x40137ffd7fff8000", "0x480280017ffc8000", "0x402580017fff8001", "0x1", "0x48127ffb7fff8000", "0x48127ffc7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480280007ffd8000", "0x480280017ffd8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff60", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x3", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff58", "0x40137fff7fff8000", "0x4003800080007ffb", "0x4003800180007ffc", "0x400380007ff97ffc", "0x402780017ff98001", "0x1", "0x4826800180008000", "0x2", "0x40297ffc7fff8002", "0x4826800180008000", "0x2", "0x480a7ffd7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff52", "0x480a7ff87fff8000", "0x480a7ffa7fff8000", "0x480680017fff8000", "0x28420862938116cb3bbdbedee07451ccc54d4e9412dbef71142ad1980a30941", "0x4829800080008002", "0x480a80007fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff67", "0x48127ffd7fff8000", "0x480a80017fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff3a", "0x40137fff7fff8000", "0x480a7ffb7fff8000", "0x480a7ffd7fff8000", "0x480680017fff8000", "0x2228d32fe428a53a1a179be3226c078dbc4ad384b11589eda25b0dbd294813b", "0x4829800080008000", "0x480a80007fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff57", "0x482480017fff8000", "0x1", "0x40307ffe7ffd7fff", "0x48127ffc7fff8000", "0x480a7ffc7fff8000", "0x480080007ffc8000", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x480680017fff8000", "0x37501df619c4fc4e96f6c0243f55e3abe7d1aca7db9af8f3740ba3696b3fdac", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffffa", "0x480a7ffb7fff8000", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff65", "0x48127ffe7fff8000", "0x48127ff57fff8000", "0x48127ff57fff8000", "0x48127ffc7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffed", "0x480a7ffa7fff8000", "0x48127ffe7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff5f", "0x48127ff67fff8000", "0x48127ff67fff8000", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x480680017fff8000", "0x1ccc09c8a19948e048de7add6929589945e25f22059c7345aaf7837188d8d05", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffffa", "0x480a7ffb7fff8000", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff47", "0x48127ffe7fff8000", "0x48127ff57fff8000", "0x48127ff57fff8000", "0x48127ffc7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffed", "0x480a7ffa7fff8000", "0x48127ffe7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff41", "0x48127ff67fff8000", "0x48127ff67fff8000", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x480680017fff8000", "0x31e7534f8ddb1628d6e07db5c743e33403b9a0b57508a93f4c49582040a2f71", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffffa", "0x480a7ffb7fff8000", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff29", "0x48127ffe7fff8000", "0x48127ff57fff8000", "0x48127ff57fff8000", "0x48127ffc7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffed", "0x480a7ffa7fff8000", "0x48127ffe7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff23", "0x48127ff67fff8000", "0x48127ff67fff8000", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x480680017fff8000", "0x13f17de67551ae34866d4aa875cbace82f3a041eaa58b1d9e34568b0d0561b", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffffa", "0x480a7ffb7fff8000", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff0b", "0x48127ffe7fff8000", "0x482480017ff78000", "0x1", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff06", "0x48127ffe7fff8000", "0x48127fee7fff8000", "0x48127fee7fff8000", "0x48127ff57fff8000", "0x48127ffb7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffe7", "0x480a7ff97fff8000", "0x48127ffe7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffeff", "0x482480017ff88000", "0x1", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffefa", "0x48127ff07fff8000", "0x48127ff07fff8000", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffebc", "0x480a7ff97fff8000", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffa7", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffc2", "0x208b7fff7fff7ffe", "0x482680017ffd8000", "0x2", "0x402a7ffd7ffc7fff", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x480280007ffd8000", "0x480280017ffd8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffec", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480280037ffb8000", "0x480680017fff8000", "0x0", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x7", "0x480a7ff57fff8000", "0x480a7ff67fff8000", "0x480a7ff87fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x2e0", "0x48127ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffebb", "0x40137fff7fff8000", "0x40137ffe7fff8001", "0x48127ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffeaf", "0x40137fff7fff8002", "0x48127ffe7fff8000", "0x48127ff07fff8000", "0x48127ff07fff8000", "0x480a7ff97fff8000", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x325", "0x40137fff7fff8003", "0x40137ffc7fff8004", "0x40137ffe7fff8005", "0x40137ffd7fff8006", "0x4829800280007ff9", "0x20680017fff7fff", "0x2e", "0x482680017ffa8000", "0x452d6860a623ebdaa6c0950cc1be6badc60b7ee699432b6ceba7d576d143119", "0x482680017ffa8000", "0x7d63192efe618520ff16b20d68c272c91616493a0787cd8abbe361b993474be", "0x48507fff7ffe8000", "0x482680017ffa8000", "0x6b8b089e4656c5f38d849f04b338559393e37995280e3277136aba00f9a2254", "0x482680017ffa8000", "0x68f0ae913141580a485fbd63911bc65ebd402c19a38d994f92b116cff2c9222", "0x48507fff7ffe8000", "0x20680017fff7ffc", "0xd", "0x480a80047fff8000", "0x480a80067fff8000", "0x480a7ff77fff8000", "0x480a80057fff8000", "0x480a80037fff8000", "0x480a80007fff8000", "0x480a80017fff8000", "0x1104800180018000", "0x2bf", "0x10780017fff7fff", "0x27", "0x20680017fff7fff", "0x15", "0x480a80047fff8000", "0x480a80067fff8000", "0x480a80057fff8000", "0x480a80007fff8000", "0x480a80017fff8000", "0x1104800180018000", "0x2eb", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480a7ff77fff8000", "0x48127ffb7fff8000", "0x480a80037fff8000", "0x48127ffa7fff8000", "0x4826800180018000", "0x1", "0x1104800180018000", "0x2c0", "0x10780017fff7fff", "0x12", "0x480a80047fff8000", "0x480a80067fff8000", "0x480a7ff77fff8000", "0x480a80057fff8000", "0x480a80037fff8000", "0x480a80007fff8000", "0x480a80017fff8000", "0x1104800180018000", "0x29f", "0x480a80037fff8000", "0x4826800180008000", "0x2", "0x4826800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffff", "0x1104800180018000", "0x2ae", "0x48127ffc7fff8000", "0x480a7ff97fff8000", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe47", "0x48127ffd7fff8000", "0x48127ff17fff8000", "0x48127ff17fff8000", "0x48127ff17fff8000", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x4003800080007ffc", "0x4826800180008000", "0x1", "0x480a7ffd7fff8000", "0x4828800080007ffe", "0x480a80007fff8000", "0x208b7fff7fff7ffe", "0x480280027ffb8000", "0x480280027ffd8000", "0x400080007ffe7fff", "0x482680017ffd8000", "0x3", "0x480280027ffd8000", "0x48307fff7ffe8000", "0x482480017fff8000", "0x1", "0x402a7ffd7ffc7fff", "0x480280027ffb8000", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280037ffb8000", "0x482480017ffc8000", "0x1", "0x480280007ffd8000", "0x480280017ffd8000", "0x480280027ffd8000", "0x482680017ffd8000", "0x3", "0x480080007ff58000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff77", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffdf", "0x48127ff37fff8000", "0x48127ff37fff8000", "0x48127ffb7fff8000", "0x48127ff27fff8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x236", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe00", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffeeb", "0x208b7fff7fff7ffe", "0x482680017ffd8000", "0x1", "0x402a7ffd7ffc7fff", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x480280007ffd8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffeb", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480280037ffb8000", "0x480680017fff8000", "0x0", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x214", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffdde", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffee7", "0x208b7fff7fff7ffe", "0x482680017ffd8000", "0x1", "0x402a7ffd7ffc7fff", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x480280007ffd8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffeb", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480280037ffb8000", "0x480680017fff8000", "0x0", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x0", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x1f0", "0x1104800180018000", "0x1f9", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffed6", "0x48307fe080007fff", "0x20680017fff7fff", "0x12", "0x48127ffa7fff8000", "0x48127ffb7fff8000", "0x48127fdd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe5b", "0x480680017fff8000", "0x1", "0x48127ffd7fff8000", "0x48307ffd80007ffe", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffdad", "0x48127ff77fff8000", "0x48127ffe7fff8000", "0x48127fdb7fff8000", "0x10780017fff7fff", "0x5", "0x48127ffa7fff8000", "0x48127ffb7fff8000", "0x48127ff97fff8000", "0x48127ffd7fff8000", "0x48127ffe7fff8000", "0x48127ffc7fff8000", "0x1104800180018000", "0x273", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe7a", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x482480017fe58000", "0x1f4", "0x48127ffb7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffec1", "0x208b7fff7fff7ffe", "0x402b7ffd7ffc7ffd", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffc6", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480280037ffb8000", "0x480680017fff8000", "0x0", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x1a9", "0x1104800180018000", "0x1b2", "0x40137fff7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe8e", "0x4828800080007fff", "0x48507ffe7fff8000", "0x20680017fff7fff", "0x7", "0x48127ff97fff8000", "0x48127ffa7fff8000", "0x48127ff87fff8000", "0x10780017fff7fff", "0xf", "0x48127ff97fff8000", "0x48127ffa7fff8000", "0x480a80007fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe0d", "0x48127ffe7fff8000", "0x482480017ffe8000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffff", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffd60", "0x48127ff87fff8000", "0x48127ffe7fff8000", "0x48127fdb7fff8000", "0x48127ffd7fff8000", "0x48127ffe7fff8000", "0x48127ffc7fff8000", "0x1104800180018000", "0x22b", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x482480017ffc8000", "0x1f4", "0x480a80007fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe7e", "0x208b7fff7fff7ffe", "0x402b7ffd7ffc7ffd", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffca", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480280037ffb8000", "0x480680017fff8000", "0x0", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x168", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe53", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffd30", "0x48127ff87fff8000", "0x48127ff87fff8000", "0x48127ff87fff8000", "0x480680017fff8000", "0x0", "0x480680017fff8000", "0x0", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe5a", "0x208b7fff7fff7ffe", "0x402b7ffd7ffc7ffd", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffe9", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480280037ffb8000", "0x480680017fff8000", "0x0", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x144", "0x1104800180018000", "0x14d", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x1104800180018000", "0x1e2", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe25", "0x48127ffd7fff8000", "0x48127ffd7fff8000", "0x48127fde7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffd09", "0x48127ff17fff8000", "0x48127ff17fff8000", "0x48127ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffddf", "0x40127fff7fff7fde", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480680017fff8000", "0x0", "0x480680017fff8000", "0x0", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffe24", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffcee", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffdf7", "0x208b7fff7fff7ffe", "0x482680017ffd8000", "0x1", "0x402a7ffd7ffc7fff", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x480280007ffd8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffcb", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480280037ffb8000", "0x480680017fff8000", "0x0", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x100", "0x1104800180018000", "0x109", "0x40137fff7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x1104800180018000", "0x19d", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffde0", "0x48127ffd7fff8000", "0x48127ffd7fff8000", "0x48127fde7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffcc4", "0x400a80007fff7ff5", "0x48127ff17fff8000", "0x48127ff17fff8000", "0x48127ffd7fff8000", "0x480680017fff8000", "0x0", "0x480680017fff8000", "0x0", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffde4", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffcae", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffd99", "0x208b7fff7fff7ffe", "0x482680017ffd8000", "0x1", "0x402a7ffd7ffc7fff", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x480280007ffd8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffcd", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480280037ffb8000", "0x480680017fff8000", "0x0", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ff77fff8000", "0x480a7ff87fff8000", "0x480a7ff97fff8000", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffd7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0xe3", "0x480a7ffb7fff8000", "0x482680017ffd8000", "0x2", "0x482680017ffc8000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffff", "0x1104800180018000", "0xf2", "0x208b7fff7fff7ffe", "0x480280027ffb8000", "0x480280017ffd8000", "0x400080007ffe7fff", "0x482680017ffd8000", "0x2", "0x480280017ffd8000", "0x48307fff7ffe8000", "0x402a7ffd7ffc7fff", "0x480280027ffb8000", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280037ffb8000", "0x482480017ffc8000", "0x1", "0x480280007ffd8000", "0x480280017ffd8000", "0x482680017ffd8000", "0x2", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffde", "0x40780017fff7fff", "0x1", "0x48127ffb7fff8000", "0x48127ffb7fff8000", "0x48127ffc7fff8000", "0x48127ffa7fff8000", "0x480680017fff8000", "0x0", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffd26", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x4003800080007ffc", "0x4826800180008000", "0x1", "0x480a7ffd7fff8000", "0x4828800080007ffe", "0x480a80007fff8000", "0x208b7fff7fff7ffe", "0x402b7ffd7ffc7ffd", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffee", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff1", "0x48127ff47fff8000", "0x48127ff47fff8000", "0x48127ffb7fff8000", "0x480280037ffb8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffd25", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x4003800080007ffc", "0x4826800180008000", "0x1", "0x480a7ffd7fff8000", "0x4828800080007ffe", "0x480a80007fff8000", "0x208b7fff7fff7ffe", "0x402b7ffd7ffc7ffd", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffee", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff1", "0x48127ff47fff8000", "0x48127ff47fff8000", "0x48127ffb7fff8000", "0x480280037ffb8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffd24", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x4003800080007ffc", "0x4826800180008000", "0x1", "0x480a7ffd7fff8000", "0x4828800080007ffe", "0x480a80007fff8000", "0x208b7fff7fff7ffe", "0x402b7ffd7ffc7ffd", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffee", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff1", "0x48127ff47fff8000", "0x48127ff47fff8000", "0x48127ffb7fff8000", "0x480280037ffb8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffd23", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x4003800080007ffb", "0x4003800180007ffc", "0x4826800180008000", "0x2", "0x480a7ffd7fff8000", "0x4828800080007ffe", "0x480a80007fff8000", "0x208b7fff7fff7ffe", "0x402b7ffd7ffc7ffd", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffed", "0x48127ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff0", "0x48127ff37fff8000", "0x48127ff37fff8000", "0x48127ffb7fff8000", "0x480280037ffb8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x302e312e30", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x4003800080007ffc", "0x4826800180008000", "0x1", "0x480a7ffd7fff8000", "0x4828800080007ffe", "0x480a80007fff8000", "0x208b7fff7fff7ffe", "0x402b7ffd7ffc7ffd", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff4", "0x480280027ffb8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff4", "0x480280007ffb8000", "0x480280017ffb8000", "0x48127ffb7fff8000", "0x480280037ffb8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbed", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbe3", "0x40127fff7fff7ff9", "0x48127ffe7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffcc1", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbbd", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffc79", "0x400a7ffd7fff7fff", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x482480017ffc8000", "0x1", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffc7e", "0x208b7fff7fff7ffe", "0x480a7ffa7fff8000", "0x482680017ffd8000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffff", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbaa", "0x480a7ff77fff8000", "0x480a7ff87fff8000", "0x48127ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffc84", "0x480a7ff97fff8000", "0x480a7ffb7fff8000", "0x48127ffd7fff8000", "0x480280007ffc8000", "0x480280017ffc8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbd9", "0x48127ff47fff8000", "0x48127ff47fff8000", "0x48127ffd7fff8000", "0x48127ff37fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x480a7ff77fff8000", "0x480a7ff87fff8000", "0x480a7ffa7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffc8f", "0x40137ffd7fff8000", "0x20680017fff7fff", "0x7", "0x48127ffc7fff8000", "0x480a80007fff8000", "0x480a7ff97fff8000", "0x48127ffb7fff8000", "0x208b7fff7fff7ffe", "0x48127ffe7fff8000", "0x482680017ffd8000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffff", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffb85", "0x48127ff77fff8000", "0x48127ffe7fff8000", "0x48127ff87fff8000", "0x480a7ffb7fff8000", "0x480a7ffd7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffc07", "0x48127ffe7fff8000", "0x480a80007fff8000", "0x480a7ff97fff8000", "0x48127ffc7fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffb57", "0x40137fff7fff8000", "0x480a80007fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffb5a", "0x482a7ffd80008000", "0x480680017fff8000", "0x657363617065", "0x400080007ffe7fff", "0x480a7ff97fff8000", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a80007fff8000", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x2", "0x480a7ff67fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffb7c", "0x40137fff7fff8000", "0x40137ffe7fff8001", "0x480a7ff77fff8000", "0x480a7ffc7fff8000", "0x480a7ffb7fff8000", "0x1104800180018000", "0x1c", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffb98", "0x48127ff77fff8000", "0x48127ffe7fff8000", "0x480a80007fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbc6", "0x480a7ff97fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbc3", "0x480a7ffa7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbc0", "0x48127fc37fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbbd", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbba", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbc8", "0x480a80017fff8000", "0x48127ffd7fff8000", "0x480a7ff87fff8000", "0x48127ffc7fff8000", "0x208b7fff7fff7ffe", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffb7e", "0x480a7ffb7fff8000", "0x48127ffe7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffb9b", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffbb9", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x480680017fff8000", "0x199d6f966f6b8e334ecb8a01c02985d6bcafad276d3dc8b25e406ca1e92d56c", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffffa", "0x480a7ffb7fff8000", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffb52", "0x48127ffe7fff8000", "0x48127ff57fff8000", "0x48127ff57fff8000", "0x48127ffc7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffed", "0x480a7ffa7fff8000", "0x48127ffe7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffb4c", "0x48127ff67fff8000", "0x48127ff67fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffe5", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x4003800080007ffc", "0x4826800180008000", "0x1", "0x480a7ffd7fff8000", "0x4828800080007ffe", "0x480a80007fff8000", "0x208b7fff7fff7ffe", "0x402b7ffd7ffc7ffd", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffee", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff1", "0x48127ff47fff8000", "0x48127ff47fff8000", "0x48127ffb7fff8000", "0x480280037ffb8000", "0x48127ffa7fff8000", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffd2", "0x208b7fff7fff7ffe", "0x482680017ffd8000", "0x1", "0x402a7ffd7ffc7fff", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x480280007ffd8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff3", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480280037ffb8000", "0x480680017fff8000", "0x0", "0x48127ffa7fff8000", "0x208b7fff7fff7ffe"], "abi": [{"inputs": [{"name": "signer", "type": "felt"}, {"name": "guardian", "type": "felt"}], "name": "constructor", "outputs": [], "type": "constructor"}, {"inputs": [{"name": "to", "type": "felt"}, {"name": "selector", "type": "felt"}, {"name": "calldata_len", "type": "felt"}, {"name": "calldata", "type": "felt*"}, {"name": "nonce", "type": "felt"}], "name": "execute", "outputs": [{"name": "response", "type": "felt"}], "type": "function"}, {"inputs": [{"name": "new_signer", "type": "felt"}], "name": "change_signer", "outputs": [], "type": "function"}, {"inputs": [{"name": "new_guardian", "type": "felt"}], "name": "change_guardian", "outputs": [], "type": "function"}, {"inputs": [], "name": "trigger_escape_guardian", "outputs": [], "type": "function"}, {"inputs": [], "name": "trigger_escape_signer", "outputs": [], "type": "function"}, {"inputs": [], "name": "cancel_escape", "outputs": [], "type": "function"}, {"inputs": [{"name": "new_guardian", "type": "felt"}], "name": "escape_guardian", "outputs": [], "type": "function"}, {"inputs": [{"name": "new_signer", "type": "felt"}], "name": "escape_signer", "outputs": [], "type": "function"}, {"inputs": [{"name": "hash", "type": "felt"}, {"name": "sig_len", "type": "felt"}, {"name": "sig", "type": "felt*"}], "name": "is_valid_signature", "outputs": [], "stateMutability": "view", "type": "function"}, {"inputs": [], "name": "get_nonce", "outputs": [{"name": "nonce", "type": "felt"}], "stateMutability": "view", "type": "function"}, {"inputs": [], "name": "get_signer", "outputs": [{"name": "signer", "type": "felt"}], "stateMutability": "view", "type": "function"}, {"inputs": [], "name": "get_guardian", "outputs": [{"name": "guardian", "type": "felt"}], "stateMutability": "view", "type": "function"}, {"inputs": [], "name": "get_escape", "outputs": [{"name": "active_at", "type": "felt"}, {"name": "caller", "type": "felt"}], "stateMutability": "view", "type": "function"}, {"inputs": [], "name": "get_version", "outputs": [{"name": "version", "type": "felt"}], "stateMutability": "view", "type": "function"}, {"inputs": [], "name": "get_block_timestamp", "outputs": [{"name": "block_timestamp", "type": "felt"}], "stateMutability": "view", "type": "function"}, {"inputs": [{"name": "new_block_timestamp", "type": "felt"}], "name": "set_block_timestamp", "outputs": [], "type": "function"}]}`

func TestProvider_Code(t *testing.T) {
	var want *types.Code
	if err := json.Unmarshal([]byte(codeDump), &want); err != nil {
		t.Fatalf("unmarshalling block: %v", err)
	}

	for _, tc := range []struct {
		address     string
		blockNumber *big.Int
	}{{
		address: "0x06eeefce63bc81620e375c7501cb7b5aecdf9fb99aa5ec25886b8b854c4293cb",
	}, {
		address:     "0x06eeefce63bc81620e375c7501cb7b5aecdf9fb99aa5ec25886b8b854c4293cb",
		blockNumber: big.NewInt(582),
	}} {
		ctx := context.Background()
		sg := NewClient(WithChain("main"))
		got, err := sg.CodeAt(ctx, tc.address, tc.blockNumber)
		if err != nil {
			t.Fatalf("getting code: %v", err)
		}

		if diff := cmp.Diff(want, got, nil); diff != "" {
			t.Errorf("Code diff mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestProviderFullContract(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
		ABIName              string
		Starknet             string
		GpsStatementVerifier string
	}
	testSet := map[string][]testSetType{
		"devnet": {},
		"mock":   {},
		"testnet": {
			{
				ABIName:              "planet_genereted",
				Starknet:             "0x04358e376b5c68f17dc1cbdbde19914f1dd6e52a2eddb5b4b0d694716fe5d89b",
				GpsStatementVerifier: "0xab43ba48c9edf4c2c4bb01237348d1d7b28ef168",
			},
		},
		"mainnet": {
			{
				ABIName:              "",
				Starknet:             "0xc662c410c0ecf747543f5ba90660f6abebd9c8c4",
				GpsStatementVerifier: "0x47312450b3ac8b5b8e247a6bb6d523e7605bdb60",
			},
		},
	}[testEnv]

	for _, test := range testSet {
		contract, err := testConfig.client.FullContract(context.Background(), test.Starknet)

		if err != nil {
			t.Fatal(err)
		}
		if contract.ABI[1].Name != test.ABIName {
			t.Fatalf("expecting %s, instead: %s", "", contract.ABI[1].Name)
		}
	}
}
