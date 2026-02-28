package domain

import "os"

// Ipfs IPFSの構造体
type Ipfs struct {
	Wallet string
	File   *os.File
}

type IpfsJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	FileType    string `json:"file_type"`
	ImageCid    string `json:"image_cid"`
	AudioCid    string `json:"audio_cid"`
	VideoCid    string `json:"video_cid"`
	Insentive   int    `json:"insentive"`
}

type IpfsAdd struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}

type IpfsPublish struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type IpfsPins struct {
	Pins     []string `json:"Pins"`
	Progress int      `json:"Progress"`
}

type IpfsResolve struct {
	Path string `json:"Path"`
}
