//go:build darwin

// macOS routes Cmd+X/C/V/A/Z through the application menu bar's key
// equivalents. webview/webview_go creates an NSApplication without any menu,
// so a WKWebView form field never receives paste no matter how focused it
// is. installEditMenu attaches a standard Edit menu to NSApp so the
// responder chain forwards those shortcuts to the focused web view.
//
// hideWindow + makeAccessory are used for the --silent re-auth path: the
// helper still needs a real WKWebView to carry the persistent cookie jar
// that login.microsoftonline.com set during interactive sign-in, but we
// don't want a window flashing on screen or a Dock icon appearing for what
// should be invisible token refresh.
package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <Cocoa/Cocoa.h>

static void hideWindow(void *nsWindow) {
    NSWindow *w = (__bridge NSWindow*)nsWindow;
    [w orderOut:nil];
}

static void makeAccessory(void) {
    [[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyProhibited];
}

static void installEditMenu(void) {
    NSApplication *app = [NSApplication sharedApplication];
    NSMenu *mainMenu = [[NSMenu alloc] init];

    NSMenuItem *appItem = [[NSMenuItem alloc] init];
    NSMenu *appMenu = [[NSMenu alloc] init];
    [appMenu addItemWithTitle:@"Hide"
                       action:@selector(hide:)
                keyEquivalent:@"h"];
    [appMenu addItem:[NSMenuItem separatorItem]];
    [appMenu addItemWithTitle:@"Quit"
                       action:@selector(terminate:)
                keyEquivalent:@"q"];
    [appItem setSubmenu:appMenu];
    [mainMenu addItem:appItem];

    NSMenuItem *editItem = [[NSMenuItem alloc] init];
    NSMenu *editMenu = [[NSMenu alloc] initWithTitle:@"Edit"];
    [editMenu addItemWithTitle:@"Undo"
                        action:@selector(undo:)
                 keyEquivalent:@"z"];
    NSMenuItem *redo = [editMenu addItemWithTitle:@"Redo"
                                           action:@selector(redo:)
                                    keyEquivalent:@"z"];
    [redo setKeyEquivalentModifierMask:(NSEventModifierFlagCommand | NSEventModifierFlagShift)];
    [editMenu addItem:[NSMenuItem separatorItem]];
    [editMenu addItemWithTitle:@"Cut"
                        action:@selector(cut:)
                 keyEquivalent:@"x"];
    [editMenu addItemWithTitle:@"Copy"
                        action:@selector(copy:)
                 keyEquivalent:@"c"];
    [editMenu addItemWithTitle:@"Paste"
                        action:@selector(paste:)
                 keyEquivalent:@"v"];
    [editMenu addItemWithTitle:@"Select All"
                        action:@selector(selectAll:)
                 keyEquivalent:@"a"];
    [editItem setSubmenu:editMenu];
    [mainMenu addItem:editItem];

    [app setMainMenu:mainMenu];
}
*/
import "C"

import "unsafe"

func installEditMenu() {
	C.installEditMenu()
}

func hideWindow(window unsafe.Pointer) {
	C.hideWindow(window)
}

func makeAccessory() {
	C.makeAccessory()
}
