package main

/*
#cgo LDFLAGS: -lgdi32
#include <Windows.h>
#include <stdio.h>
#include <stdlib.h>

#define MSG(m) {\
    MessageBoxA(NULL,m,NULL,MB_OK);}

HWND hwnd;
HWND hwndSettings;
HINSTANCE hinst;

char data[1<<30];
int PicWidth, PicHeight;

LRESULT CALLBACK WinProc(HWND hwnd,UINT msg,WPARAM wp,LPARAM lp)
{
    HDC hdc;
    PAINTSTRUCT ps;
    HPEN pen;
    static HBRUSH brush,brush2;

    switch(msg){
    case WM_DESTROY:
        PostQuitMessage(0);
        return 0;
    case WM_PAINT:
        hdc=GetDC(hwnd);

        BITMAPINFO bmi;
        memset(&bmi, 0, sizeof(bmi));
        bmi.bmiHeader.biSize = sizeof(BITMAPINFOHEADER);
        bmi.bmiHeader.biWidth = PicWidth;
        bmi.bmiHeader.biHeight = PicHeight;
        bmi.bmiHeader.biPlanes = 1;
        bmi.bmiHeader.biBitCount = 24;
        bmi.bmiHeader.biCompression = BI_RGB;
        bmi.bmiHeader.biSizeImage = 0;

        int iRet = SetDIBitsToDevice(hdc, 0, 0, PicWidth, PicHeight, 0, 0, 0, PicHeight, data, &bmi, DIB_RGB_COLORS);
        if(iRet == GDI_ERROR)
        {
            MSG("WTF!!!");
        }

        ReleaseDC(hwnd,hdc);

        return 0;
    }
    return DefWindowProc(hwnd,msg,wp,lp);
}

int WINAPI WinMain(HINSTANCE hInstance,HINSTANCE hPrevInstance,LPSTR lpCmdLine,int nShowCmd)
{
    MSG msg;
    WNDCLASS wc;

    wc.style = CS_HREDRAW | CS_VREDRAW;
    wc.lpfnWndProc=WinProc;
    wc.cbClsExtra=wc.cbWndExtra=0;
    wc.hInstance=hInstance;
    wc.hCursor=wc.hIcon=NULL;
    wc.hbrBackground=(HBRUSH)GetStockObject(BLACK_BRUSH);
    wc.lpszClassName="test";
    wc.lpszMenuName=NULL;

    if(!RegisterClass(&wc)){
        MSG("WTF!");
        return -1;
    }

    hwnd=CreateWindowA("test","testWindow",WS_VISIBLE | WS_SYSMENU | WS_CAPTION | WS_MINIMIZEBOX,0,0,PicWidth,PicHeight,NULL,NULL,hinst,NULL);

    if(hwnd==NULL){
        MSG("Window Failed.");
    }

    int check;

    while(check=GetMessage(&msg,NULL,0,0)){
        if(check==-1){
            break;
        }
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }

    UnregisterClass("test",hinst);

    return 0;

}
*/
import "C"
import (
	"image"
	_ "image/png"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("SomeFunkyPictureHere.png")
	if err != nil {
		log.Fatal(err)
	}
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	var img []byte
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			img = append(img, byte(b>>8))
			img = append(img, byte(g>>8))
			img = append(img, byte(r>>8))
		}
	}
	C.PicWidth = C.int(bounds.Max.X)
	C.PicHeight = C.int(bounds.Max.Y)
	for i, v := range img {
		C.data[i] = C.char(v)
	}
	C.WinMain(C.hinst, nil, nil, 5)
}
