#include "CustomProtocol.h"

@implementation CustomProtocolSchemeHandler
+ (void)handleGetURLEvent:(NSAppleEventDescriptor *)event withReplyEvent:(NSAppleEventDescriptor *)replyEvent {
    NSString *urlStr = [[event paramDescriptorForKeyword:keyDirectObject] stringValue];
    if (urlStr != nil) {
        HandleCustomProtocol((char*)[urlStr UTF8String]);
    }
}
@end

void StartCustomProtocolHandler(void) {
    NSAppleEventManager *appleEventManager = [NSAppleEventManager sharedAppleEventManager];

    [appleEventManager setEventHandler:[CustomProtocolSchemeHandler class]
        andSelector:@selector(handleGetURLEvent:withReplyEvent:)
        forEventClass:kInternetEventClass
        andEventID:kAEGetURL];
}
