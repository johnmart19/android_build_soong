// Copyright 2019 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

// List of VNDK libraries that have different core variant and vendor variant.
// For these libraries, the vendor variants must be installed even if the device
// has VndkUseCoreVariant set.
var VndkMustUseVendorVariantList = []string{
	"android.frameworks.sensorservice@1.0",
	"android.hardware.atrace@1.0",
	"android.hardware.audio.common@5.0",
	"android.hardware.audio.effect@2.0",
	"android.hardware.audio.effect@4.0",
	"android.hardware.audio.effect@5.0",
	"android.hardware.audio@2.0",
	"android.hardware.audio@4.0",
	"android.hardware.audio@5.0",
	"android.hardware.automotive.evs@1.0",
	"android.hardware.automotive.vehicle@2.0",
	"android.hardware.bluetooth.audio@2.0",
	"android.hardware.boot@1.0",
	"android.hardware.broadcastradio@1.0",
	"android.hardware.broadcastradio@1.1",
	"android.hardware.broadcastradio@2.0",
	"android.hardware.camera.device@1.0",
	"android.hardware.camera.device@3.2",
	"android.hardware.camera.device@3.3",
	"android.hardware.camera.device@3.4",
	"android.hardware.camera.provider@2.4",
	"android.hardware.cas.native@1.0",
	"android.hardware.cas@1.0",
	"android.hardware.configstore@1.0",
	"android.hardware.configstore@1.1",
	"android.hardware.contexthub@1.0",
	"android.hardware.drm@1.0",
	"android.hardware.drm@1.1",
	"android.hardware.fastboot@1.0",
	"android.hardware.gatekeeper@1.0",
	"android.hardware.gnss@1.0",
	"android.hardware.graphics.allocator@2.0",
	"android.hardware.graphics.bufferqueue@1.0",
	"android.hardware.graphics.composer@2.1",
	"android.hardware.graphics.composer@2.2",
	"android.hardware.health@1.0",
	"android.hardware.health@2.0",
	"android.hardware.ir@1.0",
	"android.hardware.keymaster@3.0",
	"android.hardware.keymaster@4.0",
	"android.hardware.light@2.0",
	"android.hardware.media.bufferpool@1.0",
	"android.hardware.media.omx@1.0",
	"android.hardware.memtrack@1.0",
	"android.hardware.neuralnetworks@1.0",
	"android.hardware.neuralnetworks@1.1",
	"android.hardware.neuralnetworks@1.2",
	"android.hardware.neuralnetworks@1.3",
	"android.hardware.nfc@1.0",
	"android.hardware.nfc@1.1",
	"android.hardware.nfc@1.2",
	"android.hardware.oemlock@1.0",
	"android.hardware.power.stats@1.0",
	"android.hardware.power@1.0",
	"android.hardware.power@1.1",
	"android.hardware.radio@1.4",
	"android.hardware.secure_element@1.0",
	"android.hardware.sensors@1.0",
	"android.hardware.soundtrigger@2.0",
	"android.hardware.soundtrigger@2.0-core",
	"android.hardware.soundtrigger@2.1",
	"android.hardware.tetheroffload.config@1.0",
	"android.hardware.tetheroffload.control@1.0",
	"android.hardware.thermal@1.0",
	"android.hardware.tv.cec@1.0",
	"android.hardware.tv.input@1.0",
	"android.hardware.vibrator@1.0",
	"android.hardware.vibrator@1.1",
	"android.hardware.vibrator@1.2",
	"android.hardware.weaver@1.0",
	"android.hardware.wifi.hostapd@1.0",
	"android.hardware.wifi.offload@1.0",
	"android.hardware.wifi.supplicant@1.0",
	"android.hardware.wifi.supplicant@1.1",
	"android.hardware.wifi@1.0",
	"android.hardware.wifi@1.1",
	"android.hardware.wifi@1.2",
	"android.hardwareundtrigger@2.0",
	"android.hardwareundtrigger@2.0-core",
	"android.hardwareundtrigger@2.1",
	"android.hidl.allocator@1.0",
	"android.hidl.token@1.0",
	"android.hidl.token@1.0-utils",
	"android.system.net.netd@1.0",
	"android.system.wifi.keystore@1.0",
	"libaudioroute",
	"libaudioutils",
	"libbinder",
	"libcamera_metadata",
	"libcodec2_hidl@1.0",
	"libcodec2_vndk",
	"libcrypto",
	"libdiskconfig",
	"libdumpstateutil",
	"libexpat",
	"libfmq",
	"libgatekeeper",
	"libgui",
	"libhidlcache",
	"libkeymaster_messages",
	"libkeymaster_portable",
	"libmedia_helper",
	"libmedia_omx",
	"libmemtrack",
	"libnetutils",
	"libprotobuf-cpp-full",
	"libprotobuf-cpp-lite",
	"libpuresoftkeymasterdevice",
	"libradio_metadata",
	"libselinux",
	"libsoftkeymasterdevice",
	"libsqlite",
	"libssl",
	"libstagefright_amrnb_common",
	"libstagefright_bufferpool@2.0",
	"libstagefright_bufferqueue_helper",
	"libstagefright_enc_common",
	"libstagefright_flacdec",
	"libstagefright_foundation",
	"libstagefright_omx",
	"libstagefright_omx_utils",
	"libstagefright_soft_aacdec",
	"libstagefright_soft_aacenc",
	"libstagefright_soft_amrdec",
	"libstagefright_soft_amrnbenc",
	"libstagefright_soft_amrwbenc",
	"libstagefright_soft_avcdec",
	"libstagefright_soft_avcenc",
	"libstagefright_soft_flacdec",
	"libstagefright_soft_flacenc",
	"libstagefright_soft_g711dec",
	"libstagefright_soft_gsmdec",
	"libstagefright_soft_hevcdec",
	"libstagefright_soft_mp3dec",
	"libstagefright_soft_mpeg2dec",
	"libstagefright_soft_mpeg4dec",
	"libstagefright_soft_mpeg4enc",
	"libstagefright_soft_opusdec",
	"libstagefright_soft_rawdec",
	"libstagefright_soft_vorbisdec",
	"libstagefright_soft_vpxdec",
	"libstagefright_soft_vpxenc",
	"libstagefright_softomx",
	"libstagefright_xmlparser",
	"libsysutils",
	"libtinyxml2",
	"libui",
	"libvorbisidec",
	"libxml2",
	"libyuv",
	"libziparchive",
	"vintf-vibrator-ndk_platform",
}
