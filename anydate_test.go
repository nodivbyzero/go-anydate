package goanydate

import (
	"testing"
	"time"
)

func TestAnyFormat(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "2024", want: "2006"},
		{in: "20241124", want: "20060102"},
		{in: "20241125132431", want: "20060102150405"},
		{in: "2024-11", want: "2006-01"},
		{in: "2024/11", want: "2006/01"},
		{in: "2024.11", want: "2006.01"},
		{in: "2024-11-26", want: "2006-01-02"},
		{in: "2024/11/26", want: "2006/01/02"},
		{in: "2024.11.26", want: "2006.01.02"},
		{in: "2024-9-24", want: "2006-1-02"},
		{in: "2024/9/24", want: "2006/1/02"},
		{in: "2024.9.24", want: "2006.1.02"},
		{in: "2024-9-3", want: "2006-1-2"},
		{in: "2024/9/3", want: "2006/1/2"},
		{in: "2024.9.3", want: "2006.1.2"},
		{in: "03-14-2024", want: "01-02-2006"},
		{in: "03/14/2024", want: "01/02/2006"},
		{in: "03.14.2024", want: "01.02.2006"},
		{in: "03-14-24", want: "01-02-06"},
		{in: "03/14/24", want: "01/02/06"},
		{in: "03.14.24", want: "01.02.06"},
		{in: "3-31-2024", want: "1-02-2006"},
		{in: "3/31/2024", want: "1/02/2006"},
		{in: "3.31.2024", want: "1.02.2006"},
		{in: "8-1-24", want: "1-2-06"},
		{in: "8/1/24", want: "1/2/06"},
		{in: "8.1.24", want: "1.2.06"},
		{in: "2024-11-14 22:43", want: "2006-01-02 15:04"},
		{in: "2024/11/14 22:43", want: "2006/01/02 15:04"},
		{in: "2024.11.14 22:43", want: "2006.01.02 15:04"},
		{in: "2024-11-14 22:43:57", want: "2006-01-02 15:04:05"},
		{in: "2024/11/14 22:43:57", want: "2006/01/02 15:04:05"},
		{in: "2024.11.14 22:43:57", want: "2006.01.02 15:04:05"},
		{in: "2024-11-14 10:43 AM", want: "2006-01-02 15:04 PM"},
		{in: "2024-11-14 10:43AM", want: "2006-01-02 15:04PM"},
		{in: "2024/11/14 10:43 AM", want: "2006/01/02 15:04 PM"},
		{in: "2024.11.14 10:43 AM", want: "2006.01.02 15:04 PM"},
		{in: "2024-11-14 12:43 PM", want: "2006-01-02 15:04 PM"},
		{in: "2024/11/14 12:43 PM", want: "2006/01/02 15:04 PM"},
		{in: "2024.11.14 12:43 PM", want: "2006.01.02 15:04 PM"},
		{in: "2024-11-14 10:43 am", want: "2006-01-02 15:04 pm"},
		{in: "2024/11/14 10:43 am", want: "2006/01/02 15:04 pm"},
		{in: "2024.11.14 10:43 am", want: "2006.01.02 15:04 pm"},
		{in: "2024-11-19 12:43:57 PM", want: "2006-01-02 15:04:05 PM"},
		{in: "2024/11/19 12:43:57 PM", want: "2006/01/02 15:04:05 PM"},
		{in: "2024.11.19 12:43:57 PM", want: "2006.01.02 15:04:05 PM"},
		{in: "2024-1-3 22:43", want: "2006-1-2 15:04"},
		{in: "2024/1/3 22:43", want: "2006/1/2 15:04"},
		{in: "2024.1.3 22:43", want: "2006.1.2 15:04"},
		{in: "2024-1-3 12:43 PM", want: "2006-1-2 15:04 PM"},
		{in: "2024/1/3 12:43 PM", want: "2006/1/2 15:04 PM"},
		{in: "2024.1.3 12:43 PM", want: "2006.1.2 15:04 PM"},
		{in: "2024-1-3 12:43 pm", want: "2006-1-2 15:04 pm"},
		{in: "2024/1/3 12:43 pm", want: "2006/1/2 15:04 pm"},
		{in: "2024.1.3 12:43 pm", want: "2006.1.2 15:04 pm"},
		{in: "2024-1-3 22:43:48", want: "2006-1-2 15:04:05"},
		{in: "2024/1/3 22:43:48", want: "2006/1/2 15:04:05"},
		{in: "2024.1.3 22:43:48", want: "2006.1.2 15:04:05"},
		{in: "2024-11-14 12:39:22", want: "2006-01-02 15:04:05"},
		{in: "2024/11/14 12:39:22", want: "2006/01/02 15:04:05"},
		{in: "2024.11.14 12:39:22", want: "2006.01.02 15:04:05"},
		{in: "03-14-2024 22:43", want: "01-02-2006 15:04"},
		{in: "03/14/2024 22:43", want: "01/02/2006 15:04"},
		{in: "03.14.2024 22:43", want: "01.02.2006 15:04"},
		{in: "03-14-2024 12:43 PM", want: "01-02-2006 15:04 PM"},
		{in: "03/14/2024 12:43 PM", want: "01/02/2006 15:04 PM"},
		{in: "03.14.2024 12:43 PM", want: "01.02.2006 15:04 PM"},
		{in: "03-14-2024 10:43 am", want: "01-02-2006 15:04 pm"},
		{in: "03/14/2024 10:43 am", want: "01/02/2006 15:04 pm"},
		{in: "03.14.2024 10:43 am", want: "01.02.2006 15:04 pm"},
		{in: "03-14-2024 22:43:39", want: "01-02-2006 15:04:05"},
		{in: "03/14/2024 22:43:39", want: "01/02/2006 15:04:05"},
		{in: "03.14.2024 22:43:39", want: "01.02.2006 15:04:05"},
		{in: "03-19-2024 10:43:39 AM", want: "01-02-2006 15:04:05 PM"},
		{in: "03/19/2024 10:43:39 AM", want: "01/02/2006 15:04:05 PM"},
		{in: "03.19.2024 10:43:39 AM", want: "01.02.2006 15:04:05 PM"},
		{in: "3-31-2024 22:43", want: "1-02-2006 15:04"},
		{in: "3/31/2024 22:43", want: "1/02/2006 15:04"},
		{in: "3.31.2024 22:43", want: "1.02.2006 15:04"},
		{in: "3-31-2024 22:43:41", want: "1-02-2006 15:04:05"},
		{in: "3/31/2024 22:43:41", want: "1/02/2006 15:04:05"},
		{in: "3.31.2024 22:43:41", want: "1.02.2006 15:04:05"},
		{in: "3-1-2024 22:43", want: "1-2-2006 15:04"},
		{in: "3/1/2024 22:43", want: "1/2/2006 15:04"},
		{in: "3.1.2024 22:43", want: "1.2.2006 15:04"},
		{in: "3-1-2024 22:43:42", want: "1-2-2006 15:04:05"},
		{in: "3/1/2024 22:43:42", want: "1/2/2006 15:04:05"},
		{in: "3.1.2024 22:43:42", want: "1.2.2006 15:04:05"},
		{in: "8-1-24 22:43", want: "1-2-06 15:04"},
		{in: "8/1/24 22:43", want: "1/2/06 15:04"},
		{in: "8.1.24 22:43", want: "1.2.06 15:04"},
		{in: "8-1-24 22:43:44", want: "1-2-06 15:04:05"},
		{in: "8/1/24 22:43:44", want: "1/2/06 15:04:05"},
		{in: "8.1.24 22:43:44", want: "1.2.06 15:04:05"},
		{in: "2024-11-14T12:40:15", want: "2006-01-02T15:04:05"},
		{in: "2024/11/14T12:40:15", want: "2006/01/02T15:04:05"},
		{in: "2024.11.14T12:40:15", want: "2006.01.02T15:04:05"},
		{in: "2024-11-14T13:57:23.988", want: "2006-01-02T15:04:05.999"},
		{in: "2024-11-14T13:57:23.988132", want: "2006-01-02T15:04:05.999999"},
		{in: "2024-11-14T13:57:23.988132456", want: "2006-01-02T15:04:05.999999999"},
		{in: "2024-11-14T13:57:23.900", want: "2006-01-02T15:04:05.000"},
		{in: "2024-11-14T13:57:23.988000", want: "2006-01-02T15:04:05.000000"},
		{in: "2024-11-14T13:57:23.988132000", want: "2006-01-02T15:04:05.000000000"},
		{in: "2024-11-14T16:15:09Z", want: "2006-01-02T15:04:05Z"},
		{in: "2024/11/19 16:14:52.3276369", want: "2006/01/02 15:04:05.9999999"},
		{in: "2024-Nov-19", want: "2006-Jan-02"},
		{in: "2024-November-19", want: "2006-January-02"},
		{in: "November 22, 2024", want: "January 02, 2006"},
		{in: "Nov 22, 2024", want: "Jan 02, 2006"},
		{in: "Nov 2, 2024", want: "Jan 2, 2006"},
		{in: "Nov 22, 24", want: "Jan 02, 06"},
		{in: "Nov 2, 24", want: "Jan 2, 06"},
		{in: "Nov. 22, 2024", want: "Jan. 02, 2006"},
		{in: "Nov. 22, 24", want: "Jan. 02, 06"},
		{in: "Nov. 2, 24", want: "Jan. 2, 06"},
		{in: "22 November 2024", want: "02 January 2006"},
		{in: "2 November 2024", want: "2 January 2006"},
		{in: "22 Nov 2024", want: "02 Jan 2006"},
		{in: "2 Nov 2024", want: "2 Jan 2006"},
		{in: "22 Nov 24", want: "02 Jan 06"},
		{in: "2 Nov 24", want: "2 Jan 06"},
		{in: "22/Nov/2024", want: "02/Jan/2006"},
		{in: "2/Nov/2024", want: "2/Jan/2006"},
		{in: "November 23, 2024 12:49", want: "January 02, 2006 15:04"},
		{in: "November 23, 2024 12:49:03", want: "January 02, 2006 15:04:05"},
		{in: "November 3, 2024 12:49", want: "January 2, 2006 15:04"},
		{in: "November 3, 2024 12:49:03", want: "January 2, 2006 15:04:05"},
		{in: "November 23, 2009 2:07 PM", want: "January 02, 2006 3:04 PM"},
		{in: "November 23, 2009 2:07:19 PM", want: "January 02, 2006 3:04:05 PM"},
		{in: "November 3, 2009 2:07 PM", want: "January 2, 2006 3:04 PM"},
		{in: "November 3, 2009 2:07:19 PM", want: "January 2, 2006 3:04:05 PM"},
		{in: "November 23, 2024, 12:49:03", want: "January 02, 2006, 15:04:05"},
		{in: "November 23, 2024 2:33pm", want: "January 02, 2006 3:04pm"},
		{in: "Nov 23 3:01pm", want: "Jan 02 3:04pm"},
		{in: "Nov 23 15:04:05.988", want: "Jan 02 15:04:05.999"},
		{in: "Nov 23 15:04:05.988132", want: "Jan 02 15:04:05.999999"},
		{in: "Nov 23 15:04:05.988132456", want: "Jan 02 15:04:05.999999999"},
		{in: "23 Nov 2024, 15:18", want: "02 Jan 2006, 15:04"},
		{in: "3 Nov 2024, 15:18", want: "2 Jan 2006, 15:04"},
		{in: "23 Nov 2024 16:59", want: "02 Jan 2006 15:04"},
		{in: "3 Nov 2024 16:59", want: "2 Jan 2006 15:04"},
		{in: "23 Nov 2024 17:01:40.162", want: "02 Jan 2006 15:04:05.999"},
		{in: "23/Nov/2024 17:03:05", want: "02/Jan/2006 15:04:05"},
		{in: "23.Nov.2024 17:03:05", want: "02.Jan.2006 15:04:05"},
		{in: "23-Nov-2024 17:03:05", want: "02-Jan-2006 15:04:05"},
		{in: "Sat Nov 23 17:10:06 2024", want: "Mon Jan 02 15:04:05 2006"},
		{in: "Tue Nov 5 08:04:06 2024", want: "Mon Jan 2 15:04:05 2006"},
		{in: "Tue, 26 Nov 2024 15:04:05", want: "Mon, 02 Jan 2006 15:04:05"},
		{in: "Tuesday Nov 5 08:26:06 2024", want: "Monday Jan 2 15:04:05 2006"},
		{in: "Tuesday, 26 Nov 2024 08:26:05", want: "Monday, 02 Jan 2006 15:04:05"},
		{in: "Tue 5 Nov 2024 08:29:06", want: "Mon 2 Jan 2006 15:04:05"},
		{in: "Tue 26 Nov 2024 08:29:06", want: "Mon 02 Jan 2006 15:04:05"},
		{in: "Tuesday, 26-Nov-24 08:35:05", want: "Monday, 02-Jan-06 15:04:05"},
		{in: "2024-11-26T11:48:26.371Z", want: "2006-01-02T15:04:05.999Z"},
		{in: "2024-11-26T11:45:25Z", want: "2006-01-02T15:04:05Z"},
		{in: "2024-11-26T12:08:05+0900", want: "2006-01-02T15:04:05-0700"},
		{in: "2024-11-26T12:21:05-0700", want: "2006-01-02T15:04:05-0700"},
		{in: "Tuesday, November 26 2024", want: "Monday, January 02 2006"},
		{in: "Tuesday November 26 2024", want: "Monday January 02 2006"},
		{in: "Tue, November 26 2024", want: "Mon, January 02 2006"},
		{in: "Tue November 26 2024", want: "Mon January 02 2006"},
		{in: "Sun, November 26", want: "Mon, January 02"},
		{in: "Sun November 26", want: "Mon January 02"},
		{in: "November 26", want: "January 02"},
		{in: "Nov 26", want: "Jan 02"},
		{in: "26 November", want: "02 January"},
		{in: "26 Nov", want: "02 Jan"},
		{in: "2024-11-27 09:33:51+09:00", want: "2006-01-02 15:04:05-07:00"},
		{in: "2025-11-27+09:00", want: "2006-01-02-07:00"},
		{in: "2024-11-26 10:00:43 +0700", want: "2006-01-02 15:04:05 -0700"},
		{in: "2024-11-27 10:04:44 +09:00", want: "2006-01-02 15:04:05 -07:00"},
		{in: "Wed, 27 Nov 2024 10:05:36 +0000", want: "Mon, 02 Jan 2006 15:04:05 -0700"},
		{in: "27/Nov/2024 15:04:05 -0700", want: "02/Jan/2006 15:04:05 -0700"},
		{in: "Wed, 27 Nov 2024 10:10:36 -0700", want: "Mon, 02 Jan 2006 15:04:05 -0700"},
		{in: "2024-11-27T08:15-08", want: "2006-01-02T15:04-07"},
		{in: "Wed Nov 27 10:19:05 -0700 2024", want: "Mon Jan 02 15:04:05 -0700 2006"},
		{in: "2024-11-27T10:26:09-09:00", want: "2006-01-02T15:04:05-07:00"},
		{in: "Thu Nov 28 11:37:05 MST 2024", want: "Mon Jan 02 15:04:05 MST 2006"},
		{in: "2024-11-28 12:07:00 UTC", want: "2006-01-02 15:04:05 MST"},
		{in: "Thursday, 28-Nov-24 12:11:05 MST", want: "Monday, 02-Jan-06 15:04:05 MST"},
		{in: "Thu, 28 Nov 2024 11:37:05 MST", want: "Mon, 02 Jan 2006 15:04:05 MST"},
		{in: "Thu 28 Nov 2024 12:17:09 PM UTC", want: "Mon 02 Jan 2006 15:04:05 PM MST"},
		{in: "Thursday November 28 2024 10:09am PST-08", want: "Monday January 02 2006 15:04pm MST-07"},
		{in: "Fri Nov 29 10:17:11 PST+0800 2024", want: "Mon Jan 02 15:04:05 MST-0700 2006"},
		{in: "2024-12-05 17:04:00 +0000 GMT", want: "2006-01-02 15:04:05 -0700 MST"},
		{in: "3:59PM", want: "3:04PM"},
		{in: "3-12-2024 1:12 PM", want: "1-02-2006 3:04 PM"},
		{in: "12/3/2024 1:19 PM", want: "01/2/2006 3:04 PM"},
		{in: "13:22:05.000", want: "15:04:05.000"},
		{in: "12/12/2024 13:24:59.3186369", want: "01/02/2006 15:04:05.9999999"},
		{in: "2024-12-12 13:26:59.257000000 +0000 UTC", want: "2006-01-02 15:04:05.000000000 -0700 MST"},
		{in: "Thu, 12 Dec 2024 13:29:13 +0200 (CEST)", want: "Mon, 02 Jan 2006 15:04:05 -0700 (MST)"},
		{in: "31-01-2024", want: "02-01-2006"},
		{in: "2024:12:31", want: "2006:01:02"},
		{in: "2024:12:31 16:01:51", want: "2006:01:02 15:04:05"},
		{in: "2024:12:31 16:03:51.3127369", want: "2006:01:02 15:04:05.9999999"},
		{in: "12:12:2024 16:27:09", want: "01:02:2006 15:04:05"},
		{in: "12/Dec/2024:16:36:17 -0700", want: "02/Jan/2006:15:04:05 -0700"},
		{in: "18:27:05,000", want: "15:04:05,000"},
		{in: "2024-12-14T12:49:09.99999999Z", want: "2006-01-02T15:04:05.99999999Z"},
		{in: "2024-12-14T12:59+0730", want: "2006-01-02T15:04-0700"},
		{in: "Sat, 14 Dec 2024 3:3:3 PST", want: "Mon, 02 Jan 2006 3:4:5 MST"},
	}

	for _, tt := range tests {
		got, err := DetectFormat(tt.in)
		if err != nil {
			t.Errorf("AnyFormat(\"%s\") failed with %s", tt.in, err)
		}
		if got != tt.want {
			t.Errorf("AnyFormat(\"%s\") = %s, want %s", tt.in, got, tt.want)
		}

		date, err := time.Parse(got, tt.in)
		if err != nil {
			t.Fatal(err)
		}
		if date.Format(got) != tt.in {
			t.Errorf("AnyFormat(\"%s\") failed to convert => %v", tt.in, date.Format(got))
		}
	}
}

func TestAnyFormatErr(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{

		{in: "2025-13-26"},
		{in: "2025-12-32"},
	}

	for _, tt := range tests {
		got, err := DetectFormat(tt.in)

		if err != ErrInvalidDateFormat {
			t.Errorf("AnyFormat(\"%s\") = %s, want ErrInvalidDateFormat", tt.in, got)
		}
	}
}
