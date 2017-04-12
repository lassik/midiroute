#include <CoreFoundation/CFRunLoop.h>
#include <CoreMIDI/MIDIServices.h>

#include <stdio.h>

void darwinReceiveMidiByte(int byt);

static void receiveMidiPackets(
	const MIDIPacketList *pktlist, void *refCon, void *connRefCon)
{
	const MIDIPacket *packet;
	int b, p;

	(void)refCon;
	(void)connRefCon;
	packet = pktlist->packet;
	for (p = 0; p < pktlist->numPackets; p++) {
		for (b = 0; b < packet->length; b++) {
			darwinReceiveMidiByte(packet->data[b]);
		}
		packet = MIDIPacketNext(packet);
	}
}

void darwinCoreLoop()
{
	MIDIClientRef client;
	MIDIPortRef inport;
	MIDIEndpointRef src;
	int i, n;

	MIDIClientCreate(CFSTR("MIDI Echo"), 0, 0, &client);
	MIDIInputPortCreate(
		client, CFSTR("Input port"), receiveMidiPackets, 0, &inport);
	n = MIDIGetNumberOfSources();
	fprintf(stderr, "Found %d MIDI source(s)\n", n);
	for (i = 0; i < n; i++) {
		src = MIDIGetSource(i);
		MIDIPortConnectSource(inport, src, 0);
	}
	CFRunLoopRun();
}
