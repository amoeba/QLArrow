#import <Cocoa/Cocoa.h>
#import <CoreFoundation/CoreFoundation.h>
#import <CoreServices/CoreServices.h>
#import <Foundation/Foundation.h>
#import <QuickLook/QuickLook.h>

#include "internal.h"

OSStatus GeneratePreviewForURL(void *thisInterface, QLPreviewRequestRef preview, CFURLRef url,
                               CFStringRef contentTypeUTI, CFDictionaryRef options);
void CancelPreviewGeneration(void *thisInterface, QLPreviewRequestRef preview);

/* -----------------------------------------------------------------------------
   Generate a preview for file
   This function's job is to create preview for designated file
   -----------------------------------------------------------------------------
 */
OSStatus GeneratePreviewForURL(void *thisInterface, QLPreviewRequestRef preview, CFURLRef url,
                               CFStringRef contentTypeUTI, CFDictionaryRef options) {
  struct GetParquetSummary_return ret = GetParquetSummary((char *)[[((__bridge NSURL *)url) path] cStringUsingEncoding:NSUTF8StringEncoding]);

  int err = ret.r0;
  char* data = (char*) ret.r1;
  NSString *s = [[NSString alloc] initWithCString:data encoding: NSUTF8StringEncoding];
  CFIndex len = ret.r2;

  CFDictionaryRef previewProperties = (__bridge CFDictionaryRef) @{
    (__bridge NSString *)kQLPreviewPropertyTextEncodingNameKey : @"UTF-8",
    (__bridge NSString *)kQLPreviewPropertyMIMETypeKey : @"text/html",
  };

  QLPreviewRequestSetDataRepresentation(preview, (__bridge CFDataRef)[s dataUsingEncoding:NSUTF8StringEncoding],
                                        kUTTypeHTML, previewProperties);

  return noErr;
}

void CancelPreviewGeneration(void *thisInterface, QLPreviewRequestRef preview) {
  // Implement only if supported
}
