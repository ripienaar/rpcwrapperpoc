
package main

var templates = map[string]string{
	"action": "e3sgR2VuZXJhdGVkV2FybmluZyB9fQoKcGFja2FnZSB7eyAuUGFja2FnZSB9fQoKaW1wb3J0ICgKCSJjb250ZXh0IgoJImVuY29kaW5nL2pzb24iCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tcHJvdG9jb2wvcHJvdG9jb2wiCglycGNjbGllbnQgImdpdGh1Yi5jb20vY2hvcmlhLWlvL21jb3JwYy1hZ2VudC1wcm92aWRlci9tY29ycGMvY2xpZW50IgopCgovLyB7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlcXVlc3RvciBwZXJmb3JtcyBhIFJQQyByZXF1ZXN0IHRvIHt7IC5BZ2VudE5hbWUgfCBUb0xvd2VyIH19I3t7IC5BY3Rpb25OYW1lIHwgVG9Mb3dlciB9fQp0eXBlIHt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVxdWVzdG9yIHN0cnVjdCB7CglyICAgICpyZXF1ZXN0b3IKCW91dGMgY2hhbiAqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQKfQoKLy8ge3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQgaXMgdGhlIG91dHB1dCBmcm9tIHRoZSB7eyAuQWN0aW9uTmFtZSB8IFRvTG93ZXIgfX0gYWN0aW9uCnR5cGUge3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQgc3RydWN0IHsKCWRldGFpbHMgKlJlc3VsdERldGFpbHMKCXJlcGx5ICAgbWFwW3N0cmluZ11pbnRlcmZhY2V7fQp9CgovLyB7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlc3VsdCBpcyB0aGUgcmVzdWx0IGZyb20gYSB7eyAuQWN0aW9uTmFtZSB8IFRvTG93ZXIgfX0gYWN0aW9uCnR5cGUge3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXN1bHQgc3RydWN0IHsKCXN0YXRzICAgKnJwY2NsaWVudC5TdGF0cwoJb3V0cHV0cyBbXSp7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fU91dHB1dAp9CgovLyBTdGF0cyBpcyB0aGUgcnBjIHJlcXVlc3Qgc3RhdHMKZnVuYyAoZCAqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXN1bHQpIFN0YXRzKCkgU3RhdHMgewoJcmV0dXJuIGQuc3RhdHMKfQoKLy8gUmVzdWx0RGV0YWlscyBpcyB0aGUgZGV0YWlscyBhYm91dCB0aGUgcmVxdWVzdApmdW5jIChkICp7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fU91dHB1dCkgUmVzdWx0RGV0YWlscygpICpSZXN1bHREZXRhaWxzIHsKCXJldHVybiBkLmRldGFpbHMKfQoKLy8gSGFzaE1hcCBpcyB0aGUgcmF3IG91dHB1dCBkYXRhCmZ1bmMgKGQgKnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0KSBIYXNoTWFwKCkgbWFwW3N0cmluZ11pbnRlcmZhY2V7fSB7CglyZXR1cm4gZC5yZXBseQp9CgovLyBKU09OIGlzIHRoZSBKU09OIHJlcHJlc2VudGF0aW9uIG9mIHRoZSBvdXRwdXQgZGF0YQpmdW5jIChkICp7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fU91dHB1dCkgSlNPTigpIChbXWJ5dGUsIGVycm9yKSB7CglyZXR1cm4ganNvbi5NYXJzaGFsKGQucmVwbHkpCn0KCi8vIERvIHBlcmZvcm1zIHRoZSByZXF1ZXN0CmZ1bmMgKGQgKnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVxdWVzdG9yKSBEbyhjdHggY29udGV4dC5Db250ZXh0KSAoKnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVzdWx0LCBlcnJvcikgewoJZHJlcyA6PSAme3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXN1bHR7fQoKCWhhbmRsZXIgOj0gZnVuYyhwciBwcm90b2NvbC5SZXBseSwgciAqcnBjY2xpZW50LlJQQ1JlcGx5KSB7CgkJb3V0cHV0IDo9ICZ7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fU91dHB1dHsKCQkJcmVwbHk6IG1ha2UobWFwW3N0cmluZ11pbnRlcmZhY2V7fSksCgkJCWRldGFpbHM6ICZSZXN1bHREZXRhaWxzewoJCQkJc2VuZGVyOiAgcHIuU2VuZGVySUQoKSwKCQkJCWNvZGU6ICAgIGludChyLlN0YXR1c2NvZGUpLAoJCQkJbWVzc2FnZTogci5TdGF0dXNtc2csCgkJCQl0czogICAgICBwci5UaW1lKCksCgkJCX0sCgkJfQoKCQllcnIgOj0ganNvbi5Vbm1hcnNoYWwoci5EYXRhLCAmb3V0cHV0LnJlcGx5KQoJCWlmIGVyciAhPSBuaWwgewoJCQlkLnIuY2xpZW50LmVycm9yZigiQ291bGQgbm90IGRlY29kZSByZXBseSBmcm9tICVzOiAlcyIsIHByLlNlbmRlcklEKCksIGVycikKCQl9CgoJCWlmIGQub3V0YyAhPSBuaWwgewoJCQlkLm91dGMgPC0gb3V0cHV0CgkJCXJldHVybgoJCX0KCgkJZHJlcy5vdXRwdXRzID0gYXBwZW5kKGRyZXMub3V0cHV0cywgb3V0cHV0KQoJfQoKCXJlcywgZXJyIDo9IGQuci5kbyhjdHgsIGhhbmRsZXIpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gbmlsLCBlcnIKCX0KCglkcmVzLnN0YXRzID0gcmVzCgoJcmV0dXJuIGRyZXMsIG5pbAp9CgovLyBFYWNoT3V0cHV0IGl0ZXJhdGVzIG92ZXIgYWxsIHJlc3VsdHMgcmVjZWl2ZWQKZnVuYyAoZCAqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXN1bHQpIEVhY2hPdXRwdXQoaCBmdW5jKHIgKnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0KSkgewoJZm9yIF8sIHJlc3AgOj0gcmFuZ2UgZC5vdXRwdXRzIHsKCQloKHJlc3ApCgl9Cn0KCnt7IHJhbmdlICRuYW1lLCAkaW5wdXQgOj0gLklucHV0cyB9fQovLyB7eyAkbmFtZSB8IFNuYWtlVG9DYW1lbCB9fSBpcyBhbiBpbnB1dCB0byB0aGUge3sgJC5BY3Rpb25OYW1lIHwgVG9Mb3dlciB9fSBhY3Rpb24KLy8KLy8gRGVzY3JpcHRpb246IHt7ICRpbnB1dC5EZXNjcmlwdGlvbiB9fQpmdW5jIChkICp7eyAkLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXF1ZXN0b3IpIHt7ICRuYW1lIHwgU25ha2VUb0NhbWVsIH19KHYge3sgQ2hvcmlhVHlwZVRvR29UeXBlICRpbnB1dC5UeXBlIH19KSAqe3sgJC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVxdWVzdG9yIHsKCWQuci5hcmdzWyJtZXNzYWdlIl0gPSB2CgoJcmV0dXJuIGQKfQp7eyBlbmQgfX0Ke3sgcmFuZ2UgJG5hbWUsICRvdXRwdXQgOj0gLk91dHB1dHMgfX0KLy8ge3sgJG5hbWUgfCBTbmFrZVRvQ2FtZWwgfX0gaXMgdGhlIHZhbHVlIG9mIHRoZSB7eyAkbmFtZSB9fSBvdXRwdXQKLy8KLy8gRGVzY3JpcHRpb246IHt7ICRvdXRwdXQuRGVzY3JpcHRpb24gfX0KZnVuYyAoZCAqe3sgJC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0KSB7eyAkbmFtZSB8IFNuYWtlVG9DYW1lbCB9fSgpIHt7IENob3JpYVR5cGVUb0dvVHlwZSAkb3V0cHV0LlR5cGUgfX0gewoJdmFsIDo9IGQucmVwbHlbInt7ICRuYW1lIH19Il0KCXJldHVybiB7eyBDaG9yaWFUeXBlVG9WYWxPZlR5cGUgJG91dHB1dC5UeXBlIH19Cn0Ke3sgZW5kIH19",
	"client": "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImVuY29kaW5nL2Jhc2U2NCIKCSJlbmNvZGluZy9qc29uIgoJImZtdCIKCSJ0aW1lIgoKCSJjb250ZXh0IgoKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvY2hvcmlhIgoJImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNvbmZpZyIKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1zcnZjYWNoZSIKCXJwY2NsaWVudCAiZ2l0aHViLmNvbS9jaG9yaWEtaW8vbWNvcnBjLWFnZW50LXByb3ZpZGVyL21jb3JwYy9jbGllbnQiCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vbWNvcnBjLWFnZW50LXByb3ZpZGVyL21jb3JwYy9kZGwvYWdlbnQiCgkiZ2l0aHViLmNvbS9zaXJ1cHNlbi9sb2dydXMiCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tcHJvdG9jb2wvcHJvdG9jb2wiCikKCi8vIFN0YXRzIGFyZSB0aGUgc3RhdGlzdGljcyBmb3IgYSByZXF1ZXN0CnR5cGUgU3RhdHMgaW50ZXJmYWNlIHsKCUFnZW50KCkgc3RyaW5nCglBY3Rpb24oKSBzdHJpbmcKCUFsbCgpIGJvb2wKCU5vUmVzcG9uc2VGcm9tKCkgW11zdHJpbmcKCVVuZXhwZWN0ZWRSZXNwb25zZUZyb20oKSBbXXN0cmluZwoJRGlzY292ZXJlZENvdW50KCkgaW50CglEaXNjb3ZlcmVkTm9kZXMoKSAqW11zdHJpbmcKCUZhaWxDb3VudCgpIGludAoJT0tDb3VudCgpIGludAoJUmVzcG9uc2VzQ291bnQoKSBpbnQKCVB1Ymxpc2hEdXJhdGlvbigpICh0aW1lLkR1cmF0aW9uLCBlcnJvcikKCVJlcXVlc3REdXJhdGlvbigpICh0aW1lLkR1cmF0aW9uLCBlcnJvcikKCURpc2NvdmVyeUR1cmF0aW9uKCkgKHRpbWUuRHVyYXRpb24sIGVycm9yKQp9CgovLyBOb2RlU291cmNlIGRpc2NvdmVycyBub2Rlcwp0eXBlIE5vZGVTb3VyY2UgaW50ZXJmYWNlIHsKCVJlc2V0KCkKCURpc2NvdmVyKGN0eCBjb250ZXh0LkNvbnRleHQsIGZ3IENob3JpYUZyYW1ld29yaywgZmlsdGVycyBbXUZpbHRlckZ1bmMpIChbXXN0cmluZywgZXJyb3IpCn0KCi8vIENob3JpYUZyYW1ld29yayBpcyB0aGUgY2hvcmlhIGZyYW1ld29yawp0eXBlIENob3JpYUZyYW1ld29yayBpbnRlcmZhY2UgewoJTG9nZ2VyKHN0cmluZykgKmxvZ3J1cy5FbnRyeQoJQ29uZmlndXJhdGlvbigpICpjb25maWcuQ29uZmlnCglOZXdNZXNzYWdlKHBheWxvYWQgc3RyaW5nLCBhZ2VudCBzdHJpbmcsIGNvbGxlY3RpdmUgc3RyaW5nLCBtc2dUeXBlIHN0cmluZywgcmVxdWVzdCAqY2hvcmlhLk1lc3NhZ2UpIChtc2cgKmNob3JpYS5NZXNzYWdlLCBlcnIgZXJyb3IpCglOZXdSZXBseUZyb21UcmFuc3BvcnRKU09OKHBheWxvYWQgW11ieXRlLCBza2lwdmFsaWRhdGUgYm9vbCkgKG1zZyBwcm90b2NvbC5SZXBseSwgZXJyIGVycm9yKQoJTmV3VHJhbnNwb3J0RnJvbUpTT04oZGF0YSBzdHJpbmcpIChtZXNzYWdlIHByb3RvY29sLlRyYW5zcG9ydE1lc3NhZ2UsIGVyciBlcnJvcikKCU1pZGRsZXdhcmVTZXJ2ZXJzKCkgKHNlcnZlcnMgc3J2Y2FjaGUuU2VydmVycywgZXJyIGVycm9yKQoJTmV3Q29ubmVjdG9yKGN0eCBjb250ZXh0LkNvbnRleHQsIHNlcnZlcnMgZnVuYygpIChzcnZjYWNoZS5TZXJ2ZXJzLCBlcnJvciksIG5hbWUgc3RyaW5nLCBsb2dnZXIgKmxvZ3J1cy5FbnRyeSkgKGNvbm4gY2hvcmlhLkNvbm5lY3RvciwgZXJyIGVycm9yKQoJTmV3UmVxdWVzdElEKCkgKHN0cmluZywgZXJyb3IpCglDZXJ0bmFtZSgpIHN0cmluZwp9CgovLyBGaWx0ZXJGdW5jIGNhbiBnZW5lcmF0ZSBhIGNob3JpYSBmaWx0ZXIKdHlwZSBGaWx0ZXJGdW5jIGZ1bmMoZiAqcHJvdG9jb2wuRmlsdGVyKSBlcnJvcgoKLy8ge3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHRvIHRoZSB7eyAuRERMLk1ldGFkYXRhLk5hbWUgfX0gYWdlbnQKdHlwZSB7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgc3RydWN0IHsKCWZ3ICAgICAgICAgICBDaG9yaWFGcmFtZXdvcmsKCWNmZyAgICAgICAgICAgKmNvbmZpZy5Db25maWcKCWRkbCAgICAgICAgICAgKmFnZW50LkRETAoJbnMgICAgICAgICAgICBOb2RlU291cmNlCgljbGllbnRPcHRzICAgICppbml0T3B0aW9ucwoJY2xpZW50UlBDT3B0cyBbXXJwY2NsaWVudC5SZXF1ZXN0T3B0aW9uCglmaWx0ZXJzICAgICAgIFtdRmlsdGVyRnVuYwp9CgovLyBNZXRhZGF0YSBpcyB0aGUgYWdlbnQgbWV0YWRhdGEKdHlwZSBNZXRhZGF0YSBzdHJ1Y3QgewoJTGljZW5zZSAgICAgc3RyaW5nIGBqc29uOiJsaWNlbnNlImAKCUF1dGhvciAgICAgIHN0cmluZyBganNvbjoiYXV0aG9yImAKCVRpbWVvdXQgICAgIGludCAgICBganNvbjoidGltZW91dCJgCglOYW1lICAgICAgICBzdHJpbmcgYGpzb246Im5hbWUiYAoJVmVyc2lvbiAgICAgc3RyaW5nIGBqc29uOiJ2ZXJzaW9uImAKCVVSTCAgICAgICAgIHN0cmluZyBganNvbjoidXJsImAKCURlc2NyaXB0aW9uIHN0cmluZyBganNvbjoiZGVzY3JpcHRpb24iYAp9CgovLyBNdXN0IGNyZWF0ZSBhIG5ldyBjbGllbnQgYW5kIHBhbmljcyBvbiBlcnJvcgpmdW5jIE11c3Qob3B0cyAuLi5Jbml0aWFsaXphdGlvbk9wdGlvbikgKGNsaWVudCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSB7CgljLCBlcnIgOj0gTmV3KG9wdHMuLi4pCglpZiBlcnIgIT0gbmlsIHsKCQlwYW5pYyhlcnIpCgl9CgoJcmV0dXJuIGMKfQoKLy8gTmV3IGNyZWF0ZXMgYSBuZXcgY2xpZW50IHRvIHRoZSB7eyAuRERMLk1ldGFkYXRhLk5hbWUgfX0gYWdlbnQKZnVuYyBOZXcob3B0cyAuLi5Jbml0aWFsaXphdGlvbk9wdGlvbikgKGNsaWVudCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50LCBlcnIgZXJyb3IpIHsKCWMgOj0gJnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudHsKCQlkZGw6ICAgICAgICAgICAmYWdlbnQuRERMe30sCgkJY2xpZW50UlBDT3B0czogW11ycGNjbGllbnQuUmVxdWVzdE9wdGlvbnt9LAoJCWZpbHRlcnM6ICAgICAgIFtdRmlsdGVyRnVuY3t9LAoJCWNsaWVudE9wdHM6ICZpbml0T3B0aW9uc3sKCQkJY2ZnRmlsZTogY2hvcmlhLlVzZXJDb25maWcoKSwKCQl9LAoJfQoKCWZvciBfLCBvcHQgOj0gcmFuZ2Ugb3B0cyB7CgkJb3B0KGMuY2xpZW50T3B0cykKCX0KCglpZiBjLmNsaWVudE9wdHMubnMgPT0gbmlsIHsKCQljLmNsaWVudE9wdHMubnMgPSAmQnJvYWRjYXN0TlN7fQoJfQoJYy5ucyA9IGMuY2xpZW50T3B0cy5ucwoKCWMuZncsIGVyciA9IGNob3JpYS5OZXcoYy5jbGllbnRPcHRzLmNmZ0ZpbGUpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gbmlsLCBmbXQuRXJyb3JmKCJjb3VsZCBub3QgaW5pdGlhbGl6ZSBjaG9yaWE6ICVzIiwgZXJyKQoJfQoKCWMuY2ZnID0gYy5mdy5Db25maWd1cmF0aW9uKCkKCglpZiBjLmNsaWVudE9wdHMubG9nZ2VyID09IG5pbCB7CgkJYy5jbGllbnRPcHRzLmxvZ2dlciA9IGMuZncuTG9nZ2VyKCJwdXBwZXQiKQoJfQoKCWRkbGosIGVyciA6PSBiYXNlNjQuU3RkRW5jb2RpbmcuRGVjb2RlU3RyaW5nKHJhd0RETCkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGZtdC5FcnJvcmYoImNvdWxkIG5vdCBwYXJzZSBlbWJlZGRlZCBEREw6ICVzIiwgZXJyKQoJfQoKCWVyciA9IGpzb24uVW5tYXJzaGFsKGRkbGosIGMuZGRsKQoJaWYgZXJyICE9IG5pbCB7CgkJcmV0dXJuIG5pbCwgZm10LkVycm9yZigiY291bGQgbm90IHBhcnNlIGVtYmVkZGVkIERETDogJXMiLCBlcnIpCgl9CgoJcmV0dXJuIGMsIG5pbAp9CgovLyBBZ2VudE1ldGFkYXRhIGlzIHRoZSBhZ2VudCBtZXRhZGF0YSB0aGlzIGNsaWVudCBzdXBwb3J0cwpmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIEFnZW50TWV0YWRhdGEoKSAqTWV0YWRhdGEgewoJcmV0dXJuICZNZXRhZGF0YXsKCQlMaWNlbnNlOiAgICAgcC5kZGwuTWV0YWRhdGEuTGljZW5zZSwKCQlBdXRob3I6ICAgICAgcC5kZGwuTWV0YWRhdGEuQXV0aG9yLAoJCVRpbWVvdXQ6ICAgICBwLmRkbC5NZXRhZGF0YS5UaW1lb3V0LAoJCU5hbWU6ICAgICAgICBwLmRkbC5NZXRhZGF0YS5OYW1lLAoJCVZlcnNpb246ICAgICBwLmRkbC5NZXRhZGF0YS5WZXJzaW9uLAoJCVVSTDogICAgICAgICBwLmRkbC5NZXRhZGF0YS5VUkwsCgkJRGVzY3JpcHRpb246IHAuZGRsLk1ldGFkYXRhLkRlc2NyaXB0aW9uLAoJfQp9Cgp7eyByYW5nZSAkaSwgJGFjdGlvbiA6PSAuRERMLkFjdGlvbnMgfX0KLy8ge3sgJGFjdGlvbi5OYW1lIHwgU25ha2VUb0NhbWVsIH19IHBlcmZvcm1zIHRoZSB7eyAkYWN0aW9uLk5hbWUgfCBUb0xvd2VyIH19IGFjdGlvbgovLwovLyBEZXNjcmlwdGlvbjoge3sgJGFjdGlvbi5EZXNjcmlwdGlvbiB9fQpmdW5jIChwICp7eyAkLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSB7eyAkYWN0aW9uLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX0oKSAqe3sgJGFjdGlvbi5OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVxdWVzdG9yIHsKCWQgOj0gJnt7ICRhY3Rpb24uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlcXVlc3RvcnsKCQlvdXRjOiBuaWwsCgkJcjogJnJlcXVlc3RvcnsKCQkJYXJnczogICBtYWtlKG1hcFtzdHJpbmddaW50ZXJmYWNle30pLAoJCQlhY3Rpb246ICJ7eyAkYWN0aW9uLk5hbWUgfCBUb0xvd2VyIH19IiwKCQkJY2xpZW50OiBwLAoJCX0sCgl9CgoJcmV0dXJuIGQKfQp7eyBlbmQgfX0=",
	"ddl": "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19Cgp2YXIgcmF3RERMID0gInt7IC5SYXdEREwgfCBCYXNlNjRFbmNvZGUgfX0i",
	"discover": "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImNvbnRleHQiCgkic3luYyIKCSJ0aW1lIgoKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jbGllbnQvZGlzY292ZXJ5L2Jyb2FkY2FzdCIKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1wcm90b2NvbC9wcm90b2NvbCIKKQoKdHlwZSBCcm9hZGNhc3ROUyBzdHJ1Y3QgewoJbm9kZUNhY2hlIFtdc3RyaW5nCglmICAgICAgICAgKnByb3RvY29sLkZpbHRlcgoKCXN5bmMuTXV0ZXgKfQoKZnVuYyAoYiAqQnJvYWRjYXN0TlMpIFJlc2V0KCkgewoJYi5Mb2NrKCkKCWRlZmVyIGIuVW5sb2NrKCkKCgliLm5vZGVDYWNoZSA9IFtdc3RyaW5ne30KfQoKZnVuYyAoYiAqQnJvYWRjYXN0TlMpIERpc2NvdmVyKGN0eCBjb250ZXh0LkNvbnRleHQsIGZ3IENob3JpYUZyYW1ld29yaywgZmlsdGVycyBbXUZpbHRlckZ1bmMpIChbXXN0cmluZywgZXJyb3IpIHsKCWIuTG9jaygpCglkZWZlciBiLlVubG9jaygpCgoJY29waWVyIDo9IGZ1bmMoKSBbXXN0cmluZyB7CgkJb3V0IDo9IG1ha2UoW11zdHJpbmcsIGxlbihiLm5vZGVDYWNoZSkpCgkJZm9yIGksIG4gOj0gcmFuZ2UgYi5ub2RlQ2FjaGUgewoJCQlvdXRbaV0gPSBuCgkJfQoKCQlyZXR1cm4gb3V0Cgl9CgoJaWYgIShiLm5vZGVDYWNoZSA9PSBuaWwgfHwgbGVuKGIubm9kZUNhY2hlKSA9PSAwKSB7CgkJcmV0dXJuIGNvcGllcigpLCBuaWwKCX0KCgllcnIgOj0gYi5wYXJzZUZpbHRlcnMoZmlsdGVycykKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGVycgoJfQoKCWlmIGIubm9kZUNhY2hlID09IG5pbCB7CgkJYi5ub2RlQ2FjaGUgPSBbXXN0cmluZ3t9Cgl9CgoJY2ZnIDo9IGZ3LkNvbmZpZ3VyYXRpb24oKQoJbm9kZXMsIGVyciA6PSBicm9hZGNhc3QuTmV3KGZ3KS5EaXNjb3ZlcihjdHgsIGJyb2FkY2FzdC5GaWx0ZXIoYi5mKSwgYnJvYWRjYXN0LlRpbWVvdXQodGltZS5TZWNvbmQqdGltZS5EdXJhdGlvbihjZmcuRGlzY292ZXJ5VGltZW91dCkpKQoJaWYgZXJyICE9IG5pbCB7CgkJcmV0dXJuIFtdc3RyaW5ne30sIGVycgoJfQoKCWIubm9kZUNhY2hlID0gbm9kZXMKCglyZXR1cm4gY29waWVyKCksIG5pbAp9CgpmdW5jIChiICpCcm9hZGNhc3ROUykgcGFyc2VGaWx0ZXJzKGZzIFtdRmlsdGVyRnVuYykgZXJyb3IgewoJYi5mID0gcHJvdG9jb2wuTmV3RmlsdGVyKCkKCglmb3IgXywgZiA6PSByYW5nZSBmcyB7CgkJZXJyIDo9IGYoYi5mKQoJCWlmIGVyciAhPSBuaWwgewoJCQlyZXR1cm4gZXJyCgkJfQoJfQoKCXJldHVybiBuaWwKfQo=",
	"doc": "e3sgR2VuZXJhdGVkV2FybmluZyB9fQoKLy8gUGFja2FnZSB7eyAuUGFja2FnZSB9fSBpcyBhbiBBUEkgY2xpZW50IHRvIHRoZSBDaG9yaWEge3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgQ2FwaXRhbGl6ZSB9fSBhZ2VudCBWZXJzaW9uIHt7IC5EREwuTWV0YWRhdGEuVmVyc2lvbiB9fS4KLy8KLy8gQWN0aW9uczoKe3stIHJhbmdlICRpLCAkYWN0aW9uIDo9IC5EREwuQWN0aW9ucyB9fQovLyAgICoge3sgJGFjdGlvbi5OYW1lIHwgU25ha2VUb0NhbWVsIH19IC0ge3sgJGFjdGlvbi5EZXNjcmlwdGlvbiAtfX0Ke3sgZW5kIH19CnBhY2thZ2Uge3sgLlBhY2thZ2UgfX0KCg==",
	"initoptions": "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImdpdGh1Yi5jb20vc2lydXBzZW4vbG9ncnVzIgopCgp0eXBlIGluaXRPcHRpb25zIHN0cnVjdCB7CgljZmdGaWxlIHN0cmluZwoJbG9nZ2VyICAqbG9ncnVzLkVudHJ5CglucyBOb2RlU291cmNlCn0KCi8vIEluaXRpYWxpemF0aW9uT3B0aW9uIGlzIGFuIG9wdGlvbmFsIHNldHRpbmcgdXNlZCB0byBpbml0aWFsaXplIHRoZSBjbGllbnQKdHlwZSBJbml0aWFsaXphdGlvbk9wdGlvbiBmdW5jKG9wdHMgKmluaXRPcHRpb25zKQoKLy8gQ29uZmlnRmlsZSBzZXRzIHRoZSBjb25maWcgZmlsZSB0byB1c2UsIHdoZW4gbm90IHNldCB3aWxsIHVzZSB0aGUgdXNlciBkZWZhdWx0CmZ1bmMgQ29uZmlnRmlsZShmIHN0cmluZykgSW5pdGlhbGl6YXRpb25PcHRpb24gewoJcmV0dXJuIGZ1bmMobyAqaW5pdE9wdGlvbnMpIHsKCQlvLmNmZ0ZpbGUgPSBmCgl9Cn0KCi8vIExvZ2dlciBzZXRzIHRoZSBsb2dnZXIgdG8gdXNlIGVsc2Ugb25lIGlzIG1hZGUgdmlhIHRoZSBjaG9yaWEgZnJhbWV3b3JrCmZ1bmMgTG9nZ2VyKGwgKmxvZ3J1cy5FbnRyeSkgSW5pdGlhbGl6YXRpb25PcHRpb24gewoJcmV0dXJuIGZ1bmMobyAqaW5pdE9wdGlvbnMpIHsKCQlvLmxvZ2dlciA9IGwKCX0KfQoKLy8gRGlzY292ZXJ5IHNldHMgdGhlIE5vZGVTb3VyY2UgdG8gdXNlIHdoZW4gZmluZGluZyBub2RlcyB0byBtYW5hZ2UKZnVuYyBEaXNjb3ZlcnkobnMgTm9kZVNvdXJjZSkgSW5pdGlhbGl6YXRpb25PcHRpb24gewoJcmV0dXJuIGZ1bmMobyAqaW5pdE9wdGlvbnMpIHsKCQlvLm5zID0gbnMKCX0KfQ==",
	"logging": "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgpmdW5jIChjICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIGRlYnVnZihtc2cgc3RyaW5nLCBhIC4uLmludGVyZmFjZXt9KSB7CgljLmNsaWVudE9wdHMubG9nZ2VyLkRlYnVnZihtc2csIGEuLi4pCn0KCmZ1bmMgKGMgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgaW5mb2YobXNnIHN0cmluZywgYSAuLi5pbnRlcmZhY2V7fSkgewoJYy5jbGllbnRPcHRzLmxvZ2dlci5JbmZvZihtc2csIGEuLi4pCn0KCmZ1bmMgKGMgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgd2FybmYobXNnIHN0cmluZywgYSAuLi5pbnRlcmZhY2V7fSkgewoJYy5jbGllbnRPcHRzLmxvZ2dlci5XYXJuZihtc2csIGEuLi4pCn0KCmZ1bmMgKGMgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgZXJyb3JmKG1zZyBzdHJpbmcsIGEgLi4uaW50ZXJmYWNle30pIHsKCWMuY2xpZW50T3B0cy5sb2dnZXIuRXJyb3JmKG1zZywgYS4uLikKfQo=",
	"requestor": "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImNvbnRleHQiCgkiZm10IgoKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1wcm90b2NvbC9wcm90b2NvbCIKCXJwY2NsaWVudCAiZ2l0aHViLmNvbS9jaG9yaWEtaW8vbWNvcnBjLWFnZW50LXByb3ZpZGVyL21jb3JwYy9jbGllbnQiCikKCi8vIHJlcXVlc3RvciBpcyBhIGdlbmVyaWMgcmVxdWVzdCBoYW5kbGVyCnR5cGUgcmVxdWVzdG9yIHN0cnVjdCB7CgljbGllbnQgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudAoJYWN0aW9uIHN0cmluZwoJYXJncyAgIG1hcFtzdHJpbmddaW50ZXJmYWNle30KfQoKLy8gZG8gcGVyZm9ybXMgdGhlIHJlcXVlc3QKZnVuYyAociAqcmVxdWVzdG9yKSBkbyhjdHggY29udGV4dC5Db250ZXh0LCBoYW5kbGVyIGZ1bmMocHIgcHJvdG9jb2wuUmVwbHksIHIgKnJwY2NsaWVudC5SUENSZXBseSkpICgqcnBjY2xpZW50LlN0YXRzLCBlcnJvcikgewoJci5jbGllbnQuaW5mb2YoIlN0YXJ0aW5nIGRpc2NvdmVyeSIpCgl0YXJnZXRzLCBlcnIgOj0gci5jbGllbnQubnMuRGlzY292ZXIoY3R4LCByLmNsaWVudC5mdywgci5jbGllbnQuZmlsdGVycykKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGVycgoJfQoKCWlmIGxlbih0YXJnZXRzKSA9PSAwIHsKCQlyZXR1cm4gbmlsLCBmbXQuRXJyb3JmKCJubyBub2RlcyB3ZXJlIGRpc2NvdmVyZWQiKQoJfQoJci5jbGllbnQuaW5mb2YoIkRpc2NvdmVyZWQgJWQgbm9kZXMiLCBsZW4odGFyZ2V0cykpCgoJYWdlbnQsIGVyciA6PSBycGNjbGllbnQuTmV3KHIuY2xpZW50LmZ3LCByLmNsaWVudC5kZGwuTWV0YWRhdGEuTmFtZSwgcnBjY2xpZW50LkRETChyLmNsaWVudC5kZGwpKQoJaWYgZXJyICE9IG5pbCB7CgkJcmV0dXJuIG5pbCwgZm10LkVycm9yZigiY291bGQgbm90IGNyZWF0ZSBjbGllbnQ6ICVzIiwgZXJyKQoJfQoKCW9wdHMgOj0gW11ycGNjbGllbnQuUmVxdWVzdE9wdGlvbntycGNjbGllbnQuVGFyZ2V0cyh0YXJnZXRzKX0KCWZvciBfLCBvcHQgOj0gcmFuZ2Ugci5jbGllbnQuY2xpZW50UlBDT3B0cyB7CgkJb3B0cyA9IGFwcGVuZChvcHRzLCBvcHQpCgl9CglvcHRzID0gYXBwZW5kKG9wdHMsIHJwY2NsaWVudC5SZXBseUhhbmRsZXIoaGFuZGxlcikpCgoJci5jbGllbnQuaW5mb2YoIkludm9raW5nICVzIyVzIGFjdGlvbiB3aXRoICUjdiIsIHIuY2xpZW50LmRkbC5NZXRhZGF0YS5OYW1lLCByLmFjdGlvbiwgci5hcmdzKQoKCXJlcywgZXJyIDo9IGFnZW50LkRvKGN0eCwgci5hY3Rpb24sIHIuYXJncywgb3B0cy4uLikKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGZtdC5FcnJvcmYoImNvdWxkIG5vdCBwZXJmb3JtIGRpc2FibGUgcmVxdWVzdDogJXMiLCBlcnIpCgl9CgoJcmV0dXJuIHJlcy5TdGF0cygpLCBuaWwKfQo=",
	"resultdetails": "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJInRpbWUiCgoJImdpdGh1Yi5jb20vY2hvcmlhLWlvL21jb3JwYy1hZ2VudC1wcm92aWRlci9tY29ycGMiCikKCi8vIFJlc3VsdERldGFpbHMgaXMgdGhlIGRldGFpbHMgYWJvdXQgYSByZXN1bHQKdHlwZSBSZXN1bHREZXRhaWxzIHN0cnVjdCB7CglzZW5kZXIgIHN0cmluZwoJY29kZSAgICBpbnQKCW1lc3NhZ2Ugc3RyaW5nCgl0cyAgICAgIHRpbWUuVGltZQp9CgovLyBTdGF0dXNDb2RlIGlzIGEgcmVwbHkgc3RhdHVzIGFzIGRlZmluZWQgYnkgTUNvbGxlY3RpdmUgU2ltcGxlUlBDIC0gaW50ZWdlcnMgMCB0byA1Ci8vCi8vIFNlZSB0aGUgY29uc3RhbnRzIE9LLCBSUENBYm9ydGVkLCBVbmtub3duUlBDQWN0aW9uLCBNaXNzaW5nUlBDRGF0YSwgSW52YWxpZFJQQ0RhdGEgYW5kIFVua25vd25SUENFcnJvcgp0eXBlIFN0YXR1c0NvZGUgdWludDgKCmNvbnN0ICgKCS8vIE9LIGlzIHRoZSByZXBseSBzdGF0dXMgd2hlbiBhbGwgd29ya2VkCglPSyA9IFN0YXR1c0NvZGUoaW90YSkKCgkvLyBBYm9ydGVkIGlzIHN0YXR1cyBmb3Igd2hlbiB0aGUgYWN0aW9uIGNvdWxkIG5vdCBydW4sIG1vc3QgZmFpbHVyZXMgaW4gYW4gYWN0aW9uIHNob3VsZCBzZXQgdGhpcwoJQWJvcnRlZAoKCS8vIFVua25vd25BY3Rpb24gaXMgdGhlIHN0YXR1cyBmb3IgdW5rbm93biBhY3Rpb25zIHJlcXVlc3RlZAoJVW5rbm93bkFjdGlvbgoKCS8vIE1pc3NpbmdEYXRhIGlzIHRoZSBzdGF0dXMgZm9yIG1pc3NpbmcgaW5wdXQgZGF0YQoJTWlzc2luZ0RhdGEKCgkvLyBJbnZhbGlkRGF0YSBpcyB0aGUgc3RhdHVzIGZvciBpbnZhbGlkIGlucHV0IGRhdGEKCUludmFsaWREYXRhCgoJLy8gVW5rbm93bkVycm9yIGlzIHRoZSBzdGF0dXMgZ2VuZXJhbCBmYWlsdXJlcyBpbiBhZ2VudHMgc2hvdWxkIHNldCB3aGVuIHRoaW5ncyBnbyBiYWQKCVVua25vd25FcnJvcgopCgovLyBTZW5kZXIgaXMgdGhlIGlkZW50aXR5IG9mIHRoZSByZW1vdGUgdGhhdCBwcm9kdWNlZCB0aGUgbWVzc2FnZQpmdW5jIChkICpSZXN1bHREZXRhaWxzKSBTZW5kZXIoKSBzdHJpbmcgewoJcmV0dXJuIGQuc2VuZGVyCn0KCi8vIE9LIGRldGVybWluZXMgaWYgdGhlIHJlcXVlc3Qgd2FzIHN1Y2Nlc3NmdWxsCmZ1bmMgKGQgKlJlc3VsdERldGFpbHMpIE9LKCkgYm9vbCB7CglyZXR1cm4gbWNvcnBjLlN0YXR1c0NvZGUoZC5jb2RlKSA9PSBtY29ycGMuT0sKfQoKLy8gU3RhdHVzTWVzc2FnZSBpcyB0aGUgc3RhdHVzIG1lc3NhZ2UgcHJvZHVjZWQgYnkgdGhlIHJlbW90ZQpmdW5jIChkICpSZXN1bHREZXRhaWxzKSBTdGF0dXNNZXNzYWdlKCkgc3RyaW5nIHsKCXJldHVybiBkLm1lc3NhZ2UKfQoKLy8gU3RhdHVzQ29kZSBpcyB0aGUgc3RhdHVzIGNvZGUgcHJvZHVjZWQgYnkgdGhlIHJlbW90ZQpmdW5jIChkICpSZXN1bHREZXRhaWxzKSBTdGF0dXNDb2RlKCkgU3RhdHVzQ29kZSB7CglyZXR1cm4gU3RhdHVzQ29kZShkLmNvZGUpCn0K",
	"rpcoptions": "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJInRpbWUiCgoJY29yZWNsaWVudCAiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2xpZW50L2NsaWVudCIKCXJwY2NsaWVudCAiZ2l0aHViLmNvbS9jaG9yaWEtaW8vbWNvcnBjLWFnZW50LXByb3ZpZGVyL21jb3JwYy9jbGllbnQiCikKCi8vIE9wdGlvblJlc2V0IHJlc2V0cyB0aGUgY2xpZW50IG9wdGlvbnMgdG8gdXNlIGFjcm9zcyByZXF1ZXN0cyB0byBhbiBlbXB0eSBsaXN0CmZ1bmMgKHAgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgT3B0aW9uUmVzZXQoKSAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHsKCXAuY2xpZW50UlBDT3B0cyA9IFtdcnBjY2xpZW50LlJlcXVlc3RPcHRpb257fQoJcC5ucyA9IHAuY2xpZW50T3B0cy5ucwoJcC5maWx0ZXJzID0gW11GaWx0ZXJGdW5je30KCglyZXR1cm4gcAp9CgovLyBPcHRpb25GYWN0RmlsdGVyIGFkZHMgYSBmYWN0IGZpbHRlcgpmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIE9wdGlvbkZhY3RGaWx0ZXIoZiAuLi5zdHJpbmcpICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJZm9yIF8sIGkgOj0gcmFuZ2UgZiB7CgkJcC5maWx0ZXJzID0gYXBwZW5kKHAuZmlsdGVycywgRmlsdGVyRnVuYyhjb3JlY2xpZW50LkZhY3RGaWx0ZXIoaSkpKQoJfQoKCXAubnMuUmVzZXQoKQoKCXJldHVybiBwCn0KCi8vIE9wdGlvbkNvbGxlY3RpdmUgc2V0cyB0aGUgY29sbGVjdGl2ZSB0byB0YXJnZXQKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25Db2xsZWN0aXZlKGMgc3RyaW5nKSAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHsKCXAuY2xpZW50UlBDT3B0cyA9IGFwcGVuZChwLmNsaWVudFJQQ09wdHMsIHJwY2NsaWVudC5Db2xsZWN0aXZlKGMpKQoJcmV0dXJuIHAKfQoKLy8gT3B0aW9uSW5CYXRjaGVzIHBlcmZvcm1zIHJlcXVlc3RzIGluIGJhdGNoZXMKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25JbkJhdGNoZXMoc2l6ZSBpbnQsIHNsZWVwIGludCkgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7CglwLmNsaWVudFJQQ09wdHMgPSBhcHBlbmQocC5jbGllbnRSUENPcHRzLCBycGNjbGllbnQuSW5CYXRjaGVzKHNpemUsIHNsZWVwKSkKCXJldHVybiBwCn0KCi8vIE9wdGlvbkRpc2NvdmVyeVRpbWVvdXQgY29uZmlndXJlcyB0aGUgcmVxdWVzdCBkaXNjb3ZlcnkgdGltZW91dCwgZGVmYXVsdHMgdG8gY29uZmlndXJlZCBkaXNjb3ZlcnkgdGltZW91dApmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIE9wdGlvbkRpc2NvdmVyeVRpbWVvdXQodCB0aW1lLkR1cmF0aW9uKSAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHsKCXAuY2xpZW50UlBDT3B0cyA9IGFwcGVuZChwLmNsaWVudFJQQ09wdHMsIHJwY2NsaWVudC5EaXNjb3ZlcnlUaW1lb3V0KHQpKQoJcmV0dXJuIHAKfQoKLy8gT3B0aW9uTGltaXRNZXRob2QgY29uZmlndXJlcyB0aGUgbWV0aG9kIHRvIHVzZSB3aGVuIGxpbWl0aW5nIHRhcmdldHMgLSAicmFuZG9tIiBvciAiZmlyc3QiCmZ1bmMgKHAgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgT3B0aW9uTGltaXRNZXRob2QobSBzdHJpbmcpICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJcC5jbGllbnRSUENPcHRzID0gYXBwZW5kKHAuY2xpZW50UlBDT3B0cywgcnBjY2xpZW50LkxpbWl0TWV0aG9kKG0pKQoJcmV0dXJuIHAKfQoKLy8gT3B0aW9uTGltaXRTaXplIHNldHMgbGltaXRzIG9uIHRoZSB0YXJnZXRzLCBlaXRoZXIgYSBudW1iZXIgb2YgYSBwZXJjZW50YWdlIGxpa2UgIjEwJSIKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25MaW1pdFNpemUocyBzdHJpbmcpICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJcC5jbGllbnRSUENPcHRzID0gYXBwZW5kKHAuY2xpZW50UlBDT3B0cywgcnBjY2xpZW50LkxpbWl0U2l6ZShzKSkKCXJldHVybiBwCn0KCi8vIE9wdGlvbkxpbWl0U2VlZCBzZXRzIHRoZSByYW5kb20gc2VlZCB1c2VkIHRvIHNlbGVjdCB0YXJnZXRzIHdoZW4gbGltaXRpbmcgYW5kIGxpbWl0IG1ldGhvZCBpcyAicmFuZG9tIgpmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIE9wdGlvbkxpbWl0U2VlZChzIGludDY0KSAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHsKCXAuY2xpZW50UlBDT3B0cyA9IGFwcGVuZChwLmNsaWVudFJQQ09wdHMsIHJwY2NsaWVudC5MaW1pdFNlZWQocykpCglyZXR1cm4gcAp9Cg==",
}
