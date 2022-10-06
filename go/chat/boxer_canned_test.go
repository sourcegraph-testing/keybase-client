// Copyright 2017 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// This file contains canned messages for use in testing.

package chat

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/keybase/client/go/protocol/chat1"
	"github.com/keybase/client/go/protocol/gregor1"
	"github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-codec/codec"
	"github.com/stretchr/testify/require"
)

type cannedMessage struct {
	tag                     string
	encryptionKeyGeneration int
	encryptionKey           string // hex
	verifyKey               string // hex
	senderUsername          string
	senderUID               string // hex
	senderDeviceID          string // hex
	senderDeviceName        string
	senderDeviceType        keybase1.DeviceTypeV2
	headerHash              string
	boxed                   string // hex
}

var cannedMessages = []cannedMessage{
	{
		tag: "alice25-bob25-1",
		// A message from alice25 to alice25,bob25
		// The initial tlfname message
		// Sent some time before 0148400 (pre-boxer-refactor, pre-MBv2-protocol)
		encryptionKeyGeneration: 1,
		encryptionKey:           `087655cbd18e65e4722127485d65ab70edcc7eddeddbb557e033254ebd325f32`,
		verifyKey:               `0120eacbc5f111b40a19d848c083dbff0cfb1fb43d5f041f74f48636e884d3a474a00a`,
		senderUsername:          "alice25",
		senderUID:               `5adac0bd7275ab390dd409114c7fb919`,
		senderDeviceID:          `4243553a60e774592cf0ffa508f07318`,
		senderDeviceName:        `home thing`,
		senderDeviceType:        keybase1.DeviceTypeV2_DESKTOP,
		headerHash:              `c96e286d882efc44db80e51d8e08f1de28b4934cc8491fbe88420f31106514be`,
		boxed:                   `85ae626f64794369706865727465787483a165c43884b869638e5654e0279c34cb97785f4d66bf306b7756d99ec692c95d418fc72ff7b90ede73ec4547e8697ff8c3e26d6ec46b3538cf5fa015a16ec418926965ce4c661f56a5053e0456359e6c0951402c5d15e5b1a17601ac636c69656e744865616465728aa4636f6e7683a5746c666964c410d1fec1a2287b473206e282f4d4f30116a7746f7069634944c4101fb5a5e7585a43aba1a59520939e2420a9746f7069635479706501a764656c65746573c0aa6d65726b6c65526f6f7482a468617368c4402262d87f2115b8e3ba3280ed9f535b9d6a169cec07fd50547daa2efad4a3ac837b00cbe99a0d42798a677cad8090200bc6ed0caed492630520100f73b9fe040ba57365716e6fcd60ccab6d6573736167655479706506a470726576c0a673656e646572c4105adac0bd7275ab390dd409114c7fb919ac73656e646572446576696365c4104243553a60e774592cf0ffa508f07318aa7375706572736564657300a7746c664e616d65ad616c69636532352c626f623235a9746c665075626c6963c2b06865616465724369706865727465787483a165c50177b62f90b49847ef03323e8b3f3f3007bcded0ccdae26a4d88a6cc73634f4339fbbf092fdb244d7febc42fec689ed7f1051b9999924ae20c66cc95f577d20168fc9c9a2e8999d22606de26742f9435eee2845759a071ea9f1c0cd6331970c339e30ce320fad869128766aaa7e0fba4684a7272ecac3c2f3c06d76d1884f8d316b82e1413c9ab1804e148744a078eff123f18affb51748cd1339eac62214d04d37ef4dc43b54a39de17d0d3e1f6a44ad9dc7070697cf14fa7363215c3bd9f6c76b12e74d568f0c167ab4b9df6c001e9e87e6b7731506efb06beb50a014657500f8bf94aa9507e22506901d8e15750aa39bb729d9f57cdac64bd43fd6eba81f18998a34c7b6f662d19eecb4bcd8214294123f5de58382cee42facba4aac6fd8d70f89cba1042559f5326fd4a14b8bc01caf3febddbd9c8400032ab9c8db59b229f9ee31a3ba173529d342cb5e88575bb887bee59a8fe2adb4afd567b5c1afe64dc1036e0076c9f58d02059c18004264d2b899b5408aaa98072a16ec418e752771b2921eb1e7fcdcc39f7f22f176a5518f3639b6dd5a17601ad6b657947656e65726174696f6e01ac73657276657248656164657283a56374696d65cf0000015a48b2be27a96d657373616765494401ac73757065727365646564427900`,
	},
	{
		tag: "alice25-bob25-2",
		// A normal text message
		// Sent some time before 0148400
		encryptionKeyGeneration: 1,
		encryptionKey:           `087655cbd18e65e4722127485d65ab70edcc7eddeddbb557e033254ebd325f32`,
		verifyKey:               `0120eacbc5f111b40a19d848c083dbff0cfb1fb43d5f041f74f48636e884d3a474a00a`,
		senderUsername:          "alice25",
		senderUID:               `5adac0bd7275ab390dd409114c7fb919`,
		senderDeviceID:          `4243553a60e774592cf0ffa508f07318`,
		senderDeviceName:        `home thing`,
		senderDeviceType:        keybase1.DeviceTypeV2_DESKTOP,
		headerHash:              `beb27c41dbeb2e9004f248f278243ade5e120cb7c41f40e847a2e22fe82cd3b4`,
		boxed:                   `85ae626f64794369706865727465787483a165c449ee8c699e1cb07be3b2ff1a898a06dbe202dbff79fbeb10d0fdba0cb37eb0daf039db77a237f5af9bc120537d6d973652a8bfb0b3c378f5fd3e407788d3de68eafa83b7f35408239aeea16ec4181e2cb19b55387de6f4d70ea0c2e4608e8117a8c515561f9ca17601ac636c69656e744865616465728aa4636f6e7683a5746c666964c410d1fec1a2287b473206e282f4d4f30116a7746f7069634944c4101fb5a5e7585a43aba1a59520939e2420a9746f7069635479706501a764656c65746573c0aa6d65726b6c65526f6f7482a468617368c4402262d87f2115b8e3ba3280ed9f535b9d6a169cec07fd50547daa2efad4a3ac837b00cbe99a0d42798a677cad8090200bc6ed0caed492630520100f73b9fe040ba57365716e6fcd60ccab6d6573736167655479706501a4707265769182a468617368c420c96e286d882efc44db80e51d8e08f1de28b4934cc8491fbe88420f31106514bea2696401a673656e646572c4105adac0bd7275ab390dd409114c7fb919ac73656e646572446576696365c4104243553a60e774592cf0ffa508f07318aa7375706572736564657300a7746c664e616d65ad616c69636532352c626f623235a9746c665075626c6963c2b06865616465724369706865727465787483a165c501a34974213f7ea848cf9b9b45edacc1610482d82a71b440e7ba037f4713f5e668689857c553ac52c663e74a78e75bd15788606184741d14cabcbcbeb4741b028bd27468cec658ba1c068681e1bf35566735b1124ea98fdee0633b08a52fb969b5423366d7530bb2675f60efe951e915a19ac9d1c8e95b2b40a078e9fe0de671062b3d1f780cd41794755dbdf5e1f2fe9f29fa4a578a9fc5bcb37a9f1a4e7cd55c5d7309660e0ab722b7fa1a9c499eca678e82e51c8cbf5d0d0b2f9d6e8268afc750fb2b5a6a21151db5733b5d9330d5a2dd6608e6c79cabf3bd5ae5412d844eb6fe7ce5764116dc2a5005ef407194676be1ec10a34d8b57083b808949b08d31fa74536031279ac76801922b691fd91b3ea1e7c3de8224ce1fd7e96b5560355ec67330cb3ecfc5383a9e3cb9615450a0741a0bcba2f5dada19636f8385b4d0eb75d4a0220b69dde9f0921ae58e124b8755021ac900afcedd1172c0491fb99aa6fbcf3017136298f1fb885da21ad08e55e2b583e80cdbd6c8377d53e5eb78548173360ff9d3a3b63cb6dfe799011e80c5564e93f42c120000ec0e6e66448144287d0b261b54a16ec418a5e2b86f62d26c2ed41e27b1e1c619c35aa87b4ab9e167e8a17601ad6b657947656e65726174696f6e01ac73657276657248656164657283a56374696d65cf0000015a48b2bec2a96d657373616765494402ac73757065727365646564427900`,
	},
	{
		tag: "bob25-alice25-3-deleted",
		// A TEXT message that was later edited and then deleted.
		// Sent using 0148400
		encryptionKeyGeneration: 1,
		encryptionKey:           `087655cbd18e65e4722127485d65ab70edcc7eddeddbb557e033254ebd325f32`,
		verifyKey:               `01205e1f3aa1d32555bf44f358846a141b6eecc8043e04b57bc751a9af71b314b0970a`,
		senderUsername:          "alice25",
		senderUID:               `b5b2cec12929d199a26d61938d3c3b19`,
		senderDeviceID:          `5281cf491a3a94fa897def404a7e2718`,
		senderDeviceName:        `home thing`,
		senderDeviceType:        keybase1.DeviceTypeV2_DESKTOP,
		headerHash:              `3b547a7add325ccc9f4d3012c56eb1aba01cf7687e2613493ff5c9b716afd507`,
		boxed:                   `85ae626f64794369706865727465787483a165c0a16ec0a17600ac636c69656e744865616465728aa4636f6e7683a5746c666964c410d1fec1a2287b473206e282f4d4f30116a7746f7069634944c4101fb5a5e7585a43aba1a59520939e2420a9746f7069635479706501a764656c65746573c0aa6d65726b6c65526f6f7482a468617368c4401308148fd48b7a6b7651760aa5bbfc1dbc253691d658a97e41367f3f90fec30aa88ac4d83ddc4a6b20341071f2c52bbcc5d1ac269470897568ff201427b72088a57365716e6fcd6277ab6d6573736167655479706501a4707265769182a468617368c420beb27c41dbeb2e9004f248f278243ade5e120cb7c41f40e847a2e22fe82cd3b4a2696402a673656e646572c410b5b2cec12929d199a26d61938d3c3b19ac73656e646572446576696365c4105281cf491a3a94fa897def404a7e2718aa7375706572736564657300a7746c664e616d65ad616c69636532352c626f623235a9746c665075626c6963c2b06865616465724369706865727465787483a165c501a34540f5cae386072112fd55e1b9d897625f475f3a92d429dfad6254370c914fd7f4347c0f46485a676c584c54debd7985901fbcb4e4e0d5cad5eb473f64059d4483092f3ae45158efb1fb0ed28959108fb8374691543cc003bbf91f7cc50f0b66ee1efc87e6049d07f8f52c54a148b1877923bb7bb72aa7d554ad9159a79d962cf3bbbe91b9cedd56ee0215b393de73608560c51961b201fe3a2dcd39fb688483cdf9137b7a1d256e3e66b975d6abe3f37240d475399f7ea8f154638c12c3d819203b9fe5a4f9c376e73590a7c0c15ad0b5b32503f040b86137cca52052f20e45f5d122d7864870fc9a4ce62db83745791010aad2d8a4bfee14672f2c8faeea736c3b4d2b65ba93124f4d873f7ea4acefd63aafb75b7650621fb99814b87505d05742482eec2c6f0ddd989cdaedb5f37e11051cba0f89e2125011ec516527ea6ad13a09261d87feecfc2da330aa7b729cd3843ec1150a1b5207fc3414ccbaf573befb1aeb61207b36422719172c746c3f55f44e990ec6c72c5f11ff607e437fe6a93fb15e132ae9c9954624de71b1ff8e5db0fa6bfe21f901a4060dca386ae4f1cef591a16ec418bd4c7c469e15b69d41f92ae44d7f911525609c16b27c788ca17601ad6b657947656e65726174696f6e01ac73657276657248656164657283a56374696d65cf0000015a6252df72a96d657373616765494403ac73757065727365646564427905`,
	},
	{
		tag: "bob25-alice25-4-deleted-edit",
		// An EDIT of the previous message. Then deleted by the next message.
		// Sent using 0148400
		encryptionKeyGeneration: 1,
		encryptionKey:           `087655cbd18e65e4722127485d65ab70edcc7eddeddbb557e033254ebd325f32`,
		verifyKey:               `01205e1f3aa1d32555bf44f358846a141b6eecc8043e04b57bc751a9af71b314b0970a`,
		senderUsername:          "bob25",
		senderUID:               `b5b2cec12929d199a26d61938d3c3b19`,
		senderDeviceID:          `5281cf491a3a94fa897def404a7e2718`,
		senderDeviceName:        `home thing`,
		senderDeviceType:        keybase1.DeviceTypeV2_DESKTOP,
		headerHash:              `ea685e0f26b5b4fc1de415113440cc3d5465a15242d683a7f48896ecd2c6d626`,
		boxed:                   `85ae626f64794369706865727465787483a165c0a16ec0a17600ac636c69656e744865616465728ca4636f6e7683a5746c666964c410d1fec1a2287b473206e282f4d4f30116a7746f7069634944c4101fb5a5e7585a43aba1a59520939e2420a9746f7069635479706501a764656c65746573c0aa6d65726b6c65526f6f7482a468617368c4401308148fd48b7a6b7651760aa5bbfc1dbc253691d658a97e41367f3f90fec30aa88ac4d83ddc4a6b20341071f2c52bbcc5d1ac269470897568ff201427b72088a57365716e6fcd6277ab6d6573736167655479706503a86f7574626f784944c4088ecc94b7ff505c04aa6f7574626f78496e666f82ab636f6d706f736554696d65cf0000015a62544240a47072657603a4707265769182a468617368c4203b547a7add325ccc9f4d3012c56eb1aba01cf7687e2613493ff5c9b716afd507a2696403a673656e646572c410b5b2cec12929d199a26d61938d3c3b19ac73656e646572446576696365c4105281cf491a3a94fa897def404a7e2718aa7375706572736564657303a7746c664e616d65ad616c69636532352c626f623235a9746c665075626c6963c2b06865616465724369706865727465787483a165c501ddb022060f0a3fa525740ad805de68991945b7a973700d7f18f93a23bff9d6d7377f4b05308737f0220fc0a6ad826bb3e75fde81ce8f4af0408797a80a38a6655b754d3f03f1bc313c426e81a5ae3ce10557c3d5bfae769ee811535c10f58be3ff9fff06f9a1135d23cf093e4df78d096a53a7fc062f5944c32baf7af0a00cb26fb522378db88e275ea2d8824fe049cc1b1fee5e7f204baf0f6348ac43e7248ab7011b722b39fe0ba025b057fc72637b6257bd3d40bdd84a3d9d697365382bfa7a3b70f22d40766d992b366a25c19103e342e514dc04aefa84fd638bbc9faec91895bba04f5f232301624f24125547a528db3eaef7d28ce0c0df70f4aa8d0c15cf449bca18efae03bacdbea06f9c664dac0f2e40d56b77df0e0ed27fd127fc92312d5acd1d7bbb024ec85c2a41909fb7424b806933b639867e5724cb93ba89f7cae027640c2ebe6e2eacb1e084d32fb428907ada22677b4bd23429ffccfc90da58487c752b1c14a21aab58589b9b3656df3478a71a4ff56838ad0afabf761f86161209ab0a5f33c4be212087443ad84fc210f8c2939ac1f95fa71e4679a502bab6cb72b665b054181873a673abc5fa29d17caec7ed4bc7217375145fda6d25236307829b002f23378fb7439ae4d5ce68548448a5ca5df8de0c0dd3db9ebaa16ec4185c9dcf2b658a4f44174e6f21a887b6053009de63fce942a3a17601ad6b657947656e65726174696f6e01ac73657276657248656164657283a56374696d65cf0000015a62544252a96d657373616765494404ac73757065727365646564427900`,
	},
	{
		tag: "bob25-alice25-5-delete",
		// A DELETE that deletes a TEXT with one EDIT
		// Sent using 0148400
		encryptionKeyGeneration: 1,
		encryptionKey:           `087655cbd18e65e4722127485d65ab70edcc7eddeddbb557e033254ebd325f32`,
		verifyKey:               `01205e1f3aa1d32555bf44f358846a141b6eecc8043e04b57bc751a9af71b314b0970a`,
		senderUsername:          "bob25",
		senderUID:               `b5b2cec12929d199a26d61938d3c3b19`,
		senderDeviceID:          `5281cf491a3a94fa897def404a7e2718`,
		senderDeviceName:        `home thing`,
		senderDeviceType:        keybase1.DeviceTypeV2_DESKTOP,
		headerHash:              `68093cfb9f2a4ab1d872c9459a4796817cbf7f944ba220b10101efb4f84ad518`,
		boxed:                   `85ae626f64794369706865727465787483a165c44e90846a44fea43b1186dbe5faa73bc6ba57f68d22c74c8fb3c8e5187cf4a13c261910017c90236f61210bc2630636bacf712aff0f2f7c0a3d5f6c36f7542523d75f50926453b62f7f63918777e0a3a16ec4186d4cd8b8eb35a53e73795b57543af0cb63dec6177d4eea46a17601ac636c69656e744865616465728ca4636f6e7683a5746c666964c410d1fec1a2287b473206e282f4d4f30116a7746f7069634944c4101fb5a5e7585a43aba1a59520939e2420a9746f7069635479706501a764656c65746573920304aa6d65726b6c65526f6f7482a468617368c4401308148fd48b7a6b7651760aa5bbfc1dbc253691d658a97e41367f3f90fec30aa88ac4d83ddc4a6b20341071f2c52bbcc5d1ac269470897568ff201427b72088a57365716e6fcd6277ab6d6573736167655479706504a86f7574626f784944c408dc74065df95f1c48aa6f7574626f78496e666f82ab636f6d706f736554696d65cf0000015a62546d28a47072657603a4707265769182a468617368c420ea685e0f26b5b4fc1de415113440cc3d5465a15242d683a7f48896ecd2c6d626a2696404a673656e646572c410b5b2cec12929d199a26d61938d3c3b19ac73656e646572446576696365c4105281cf491a3a94fa897def404a7e2718aa7375706572736564657303a7746c664e616d65ad616c69636532352c626f623235a9746c665075626c6963c2b06865616465724369706865727465787483a165c501dde31898b61cbcdf296e28a48076ed58977d7134ca7093a8c067a4eb9a82f7aad5399bec78d30a5c38d6369dc48b393246fe137cb54389903d074e55526cfc7e5ff0dbbdf30c03e4b5542f080e0ea2b9c2ca39b113ed1aef5765b41d2bbc681967671ade1b59a0333cc321f46c380f550edfc42979bda6b418d1c74e8ca7e875fcd8dd9f1bc1a73e2c22dd9d031a23921188a498f73930569b987141868d0e9f299abb58f35a53a6d0ff3dddfaa00b7cc8f6b1d358555dfe675fc8416704969ab96aa87aa3ccddbb0c6d342dc7478dfbf4440b4456fc5937c1f247366b9c2aa4659197314012f460ebb44aec4f84ef60bb2d1437967a9bb25c5a0a05bc98c85fcdee14a908d1976133153d1cbc6e96b83d18df0c9b3be78d13cdb2ad7eb01b3fb5e8f5831f98cf76c435b55634459b731058552598c6d8690183355326367e598e7739b595b7c03c5fe0ed86cff70340896a53ceee49c6fc2a5e9f9d78c714374e6147e9b2e76e32fdd07f4ff9c94b816537a8b16d0a5f500bb04f13ca910f15e5726ece786f97e144e358b410371ae1d022ce97ef37d3c21b75500a0467b575bcb1b144c7e4646eb5bd80042c997941faa81c21d5aa38d6b645cdbb0a7b05ae3141e5452d83095f7bee2f2236813f22acd8a9f89fdd30cc1b561065907aa16ec418b99852fbd55705e0c4c2195965f0c3ef21ffa631feae79d7a17601ad6b657947656e65726174696f6e01ac73657276657248656164657283a56374696d65cf0000015a62546d47a96d657373616765494405ac73757065727365646564427900`,
	},
}

func getCannedMessage(t *testing.T, tag string) cannedMessage {
	for _, cm := range cannedMessages {
		if cm.tag == tag {
			return cm
		}
	}
	errStr := fmt.Sprintf("Cannot find canned message: %q", tag)
	t.Fatalf(errStr)
	panic(errStr)
}

func (cm cannedMessage) AsBoxed(t *testing.T) (res chat1.MessageBoxed) {
	cm.unhex(t, &res, cm.boxed)
	return res
}

func (cm cannedMessage) EncryptionKey(t *testing.T) *keybase1.CryptKey {
	keyBytes, err := hex.DecodeString(cm.encryptionKey)
	require.NoError(t, err)

	var keyBytes32 keybase1.Bytes32
	copy(keyBytes32[:], keyBytes)
	res := keybase1.CryptKey{
		KeyGeneration: cm.encryptionKeyGeneration,
		Key:           keyBytes32,
	}
	require.NotNil(t, res, "nil canned encryption key")
	require.Equal(t, len(res.Key), 32, "non-32 byte encryption key")
	return &res
}

func (cm cannedMessage) VerifyKey(t *testing.T) []byte {
	keyBytes, err := hex.DecodeString(cm.verifyKey)
	require.NoError(t, err)
	require.NotNil(t, keyBytes)
	return keyBytes
}

func (cm cannedMessage) SenderUID(t *testing.T) gregor1.UID {
	uid, err := hex.DecodeString(cm.senderUID)
	require.NoError(t, err)
	require.NotNil(t, uid)
	return uid
}

func (cm cannedMessage) SenderDeviceID(t *testing.T) gregor1.DeviceID {
	did, err := hex.DecodeString(cm.senderDeviceID)
	require.NoError(t, err)
	require.NotNil(t, did)
	return did
}

func (cm cannedMessage) unhex(t *testing.T, out any, inHex string) {
	bytes, err := hex.DecodeString(inHex)
	require.NoError(t, err)
	mh := codec.MsgpackHandle{WriteExt: true}
	dec := codec.NewDecoderBytes(bytes, &mh)
	err = dec.Decode(&out)
	require.NoError(t, err)
}
