package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDict(t *testing.T) {
	cases := []struct {
		input  string
		expect Doc
	}{
		{
			input:  "",
			expect: Doc{},
		},
		{
			input: `<?xml version="1.0" encoding="UTF-8"?>
<xml leorendertarget="1" dictQueryXSLT="0" lion="0" api="" lp="ende" lang="de">
    <advMedia url="/advMedia/ende-74b6fc66.xml"/>
    <sectionlist sectionsort="bestPrio">
        <section sctnum="2" sctName="adjadv" sctTitle="Adjektive/Adverbien" sctCount="16" sctTotalCnt="156" sctDirectCnt="23">
            <entry uid="263280" langlvl="B">
                <side hc="1" lang="en">
                    <words>
                        <word>short</word>
                    </words>
                </side>
                <side hc="0" lang="de">
                    <words>
                        <word>kurz</word>
                    </words>
                </side>
            </entry>
        </section>
    </sectionlist>
</xml>`,
			expect: Doc{
				Sections: []Section{
					{
						SectionName: "adjadv",
						Entries: []Entry{
							{
								Sides: []Side{
									{
										Lang: "en",
										Word: "short",
									}, {
										Lang: "de",
										Word: "kurz",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range cases {
		got, err := ProcessQueryXml(bytes.NewBufferString(test.input))
		if err != nil {
			t.Fatalf("ProcessQueryXml failed for input: %s", test.input)
		}
		if !reflect.DeepEqual(got, test.expect) {
			t.Fatalf("Expected: %s, got: %s", test.expect, got)
		}
	}
}
