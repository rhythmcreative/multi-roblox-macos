#import "app.h"

extern void goAppClean(void);
extern void goCloseAll(void);

@interface MultiRobloxAppDelegate : NSObject <NSApplicationDelegate, NSWindowDelegate>
@property (strong, nonatomic) NSWindow *window;
@end

@implementation MultiRobloxAppDelegate

- (void)applicationDidFinishLaunching:(NSNotification *)notification {
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
    [NSApp activateIgnoringOtherApps:YES];

    NSRect frame = NSMakeRect(0, 0, 320, 160);
    NSUInteger style = NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskMiniaturizable;
    self.window = [[NSWindow alloc] initWithContentRect:frame
                                              styleMask:style
                                                backing:NSBackingStoreBuffered
                                                  defer:NO];
    [self.window setTitle:@"Multi Roblox macOS"];
    [self.window center];
    [self.window setDelegate:self];

    NSTextField *label = [NSTextField labelWithString:@"Multi-Roblox is active!\nLaunch games from Roblox website."];
    [label setAlignment:NSTextAlignmentCenter];

    NSButton *gitBtn = [NSButton buttonWithTitle:@"Fork on GitHub" target:self action:@selector(openGitHub:)];
    NSButton *closeBtn = [NSButton buttonWithTitle:@"Close all Roblox instances" target:self action:@selector(closeRoblox:)];

    NSStackView *stack = [NSStackView stackViewWithViews:@[label, gitBtn, closeBtn]];
    [stack setOrientation:NSUserInterfaceLayoutOrientationVertical];
    [stack setDistribution:NSStackViewDistributionFillEqually];
    [stack setAlignment:NSLayoutAttributeCenterX];
    [stack setSpacing:12];

    [self.window.contentView addSubview:stack];
    stack.translatesAutoresizingMaskIntoConstraints = NO;
    [NSLayoutConstraint activateConstraints:@[
        [stack.topAnchor constraintEqualToAnchor:self.window.contentView.topAnchor constant:16],
        [stack.bottomAnchor constraintEqualToAnchor:self.window.contentView.bottomAnchor constant:-16],
        [stack.leadingAnchor constraintEqualToAnchor:self.window.contentView.leadingAnchor constant:20],
        [stack.trailingAnchor constraintEqualToAnchor:self.window.contentView.trailingAnchor constant:-20]
    ]];

    [self.window makeKeyAndOrderFront:nil];
}

- (void)openGitHub:(id)sender {
    [[NSWorkspace sharedWorkspace] openURL:[NSURL URLWithString:@"https://github.com/rhythmcreative/multi-roblox-macos"]];
}

- (void)closeRoblox:(id)sender {
    goCloseAll();
}

- (BOOL)windowShouldClose:(NSWindow *)sender {
    [NSApp terminate:nil];
    return YES;
}

- (void)applicationWillTerminate:(NSNotification *)notification {
    goAppClean();
}

@end

void RunNativeApp(void) {
    @autoreleasepool {
        NSApplication *app = [NSApplication sharedApplication];
        MultiRobloxAppDelegate *delegate = [[MultiRobloxAppDelegate alloc] init];
        [app setDelegate:delegate];
        [app run];
    }
}
