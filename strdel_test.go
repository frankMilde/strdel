package strdel

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/frankMilde/rol/testutils"
)

func Test_TrailingSpaces_haveSpacesBeforeLinebreaks_spaceAreRemoved(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `a 
b			
c     


d




e   	 			 	 	





`,
			Want: `a
b
c


d




e





`,
		},
	}

	for _, test := range tests {
		got := TrailingSpaces(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall TrailingSpaces(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_LeadingSpaces_haveSpacesBeforeFireChar_spaceAreRemoved(t *testing.T) {
	tests := testutils.ConversionTests{
		{ // simple spaces
			In:   `   a `,
			Want: `a `,
		},
		{ // tabs
			In: `		a	`,
			Want: `a	`,
		},
		{ // newline spaces
			In: `
    <li>Services\EntityService.cs</li>
    <li>Services\GameService.cs</li>
    <li>Services\IEntityService.cs</li>
    <li>Services\PlayerService.cs</li>

`,
			Want: `
<li>Services\EntityService.cs</li>
<li>Services\GameService.cs</li>
<li>Services\IEntityService.cs</li>
<li>Services\PlayerService.cs</li>

`,
		},
		{ // newline tabs
			In: `
	<li>Services\EntityService.cs</li>
 			<li>Services\GameService.cs</li>
 <li>Services\IEntityService.cs</li>
						<li>Services\PlayerService.cs</li>

`,
			Want: `
<li>Services\EntityService.cs</li>
<li>Services\GameService.cs</li>
<li>Services\IEntityService.cs</li>
<li>Services\PlayerService.cs</li>

`,
		},
	}

	for _, test := range tests {
		got := LeadingSpaces(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall LeadingSpaces(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_EmptyBrackets_haveMultilinedEmptyBrackets_spaceAndNewlinesAreRemoved(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `\emph{ 


			}`,
			Want: `\emph{}`,
		},
		{
			In:   `\emph{\\}`,
			Want: `\emph{}`,
		},
	}

	for _, test := range tests {
		got := EmptyBrackets(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall EmptyBrackets(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_EmptyLines_haveEmptyMultiLine_RemoveEmptyLines(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `\emph{
 
		
			}`,
			Want: `\emph{
			}`,
		},
	}

	for _, test := range tests {
		got := EmptyLine(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall EmptyLine(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_EmptyNestedMacros_haveEmptyNestedMacros_AreRemoved(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In:   `test \underline{} test`,
			Want: `test  test`,
		},
		{
			In:   `test \underline{\textbf{}} test`,
			Want: `test  test`,
		},
		{
			In:   `test \underline{test} test`,
			Want: `test \underline{test} test`,
		},
		{
			In:   `test \underline{\textbf{test}} test`,
			Want: `test \underline{\textbf{test}} test`,
		},
	}

	for _, test := range tests {
		got := EmptyMacros(test.In, 2)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall EmptyBrackets(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_SpaceBeforeClosingBrackets_haveMultilinedSpaceBeforeBrackets_spaceAndNewlinesAreRemoved(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `Test.
			Test. \emph{Test  }
			
			Test`,
			Want: `Test.
			Test. \emph{Test} 
			
			Test`,
		},
		{
			In: `Test.
			Test. \emph{Test


			}
			
			Test`,
			Want: `Test.
			Test. \emph{Test} 
			
			Test`,
		},
	}

	for _, test := range tests {
		got := SpaceBeforeClosingBrackets(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall SpaceBeforeClosingBrackets(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}
func Test_SpaceBeforeClosingBrackets_hasNoSpaceBeforeAndAfterBrackets_NoNewSpaceShouldBeAdded(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `Test.
			Test. \emph{Test}Test
			
			Test`,
			Want: `Test.
			Test. \emph{Test}Test
			
			Test`,
		},
	}

	for _, test := range tests {
		got := SpaceBeforeClosingBrackets(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall SpaceBeforeClosingBrackets(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}
func Test_SpaceBeforeClosingBrackets_haveTexLinebreakBeforeBrackets_spaceAndNewlinesAreRemoved(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `Test1.
			Test. \emph{Test1 \\}
			
			Test`,
			Want: `Test1.
			Test. \emph{Test1} \\
			
			Test`,
		},
		{
			In: `Test2.
			Test. \emph{Test2 \\ \\ }
			
			Test`,
			Want: `Test2.
			Test. \emph{Test2} \\ 
			
			Test`,
		},
		{
			In: `Test3.
			Test. \emph{Test3


			\\}
			
			Test`,
			Want: `Test3.
			Test. \emph{Test3} \\
			
			Test`,
		},
	}

	for _, test := range tests {
		got := SpaceBeforeClosingBrackets(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall SpaceBeforeClosingBrackets(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_EmptyLinesInMacros_haveEmptyNewLines_NoNewLinesInMacro(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `Test
			 \emph{ Test1 Test1
		
		Test1 }

			Test
			`,
			Want: `Test
			 \emph{ Test1 Test1 
			 Test1 }

			Test
			`,
		},
		{
			In: `
\begin{Figure}
        \centering
        \includegraphics[width=0.95\textwidth]{/home/frank/rol/615016607.jpeg}
        \captionof{figure}{Photographer: David Paul Morris/Bloomberg \href{https://assets.bwbx.io/images/users/iqjWHBFdfxIU/iCbVQnix8bsQ/v1/640x-1.jpg}{[Image Source]}}
        \label{fig:615016607}
\end{Figure}

        \href{https://subscribe.businessweek.com/servlet/OrdersGateway?cds_mag_code=BWK&cds_page_id=205566}{}

In the world of artificial intelligence, one of the year's biggest coming-out parties is the Neural Information Processing Systems conference. Thousands of researchers from universities and software companies gather to share their work and wrestle with new ways to tailor software to people's habits. At last year's conference in Montreal, employees of Google, Microsoft, and IBM presented papers on teaching computers to work faster and smarter, such as by reading the house numbers in a photo to determine an address. But one player was conspicuously absent: Apple. This year, Chinese search giant Baidu and Facebook, along with Google and Microsoft, are slated to present papers. Apple isn't.

Apple researchers attended the Montreal conference last year but kept a low profile and didn't say who they worked for unless asked, says Yoshua Bengio, an AI pioneer and professor of computer science at the University of Montreal. This is typical of the company's appearances at the field's big AI conferences, say Bengio and other prominent researchers. “Apple is off the scale in terms of secrecy,” says Richard Zemel, a
professor in the computer science department at the University of Toronto. “They're completely out of the loop.” Apple declined to comment for this story.

Other big consumer-software companies have set up \href{http://www.bloomberg.com/news/features/2015-06-09/the-future-of-computers-is-the-mind-of-a-toddler}{research centers} staffed with dozens or hundreds of AI
experts from around the world, racing to publish findings. The fast-paced culture has produced Google algorithms that can more accurately transcribe speech and tag photos; personal assistants that are smarter than Siri, such as Microsoft's \href{http://www.bloomberg.com/news/videos/b/c2986a39-ede9-4907-8226-fa28a6a57086}{Cortana} and Google's \href{http://www.bloomberg.com/news/articles/2015-05-28/what-google-just-announced-is-a-bombshell}{Now}; and Facebook software that can tell blind users what's in their friends' photos.
\begin{quotation}

 \textbf{0}

Number of AI papers Apple researchers have published
\end{quotation}

At Apple, new hires on the AI teams are told not to announce their positions on LinkedIn or Twitter, says Graham Taylor, a professor of machine learning at Ontario's University of Guelph. They've been told to lock their office doors whenever they leave, says a person familiar with the matter. The teams working to make the company's software smarter are kept unaware of what similar Apple teams are doing, say two people who've worked with the company.

So far, \href{http://www.bloomberg.com/news/videos/2014-09-25/apple-iphone-bends-crashes-into-software-road-bump}{so Apple}. For years, the fortunes of the world's most valuable company have been associated with
its opacity. But its biggest AI success to date has been buying Siri from a \href{http://www.bloomberg.com/bw/articles/2013-02-14/siri-creator-sri-spins-off-a-new-virtual-assistant}{startup} in 2010. Apple Maps still lags the predictive capabilities of similar software. “There's no way they can just observe and not be part of the community and take advantage of what is going on,” Bengio says. “I believe if they don't change their attitude, they will stay behind.”

What makes developing AI different from a mobile operating systemapart from the uncharted technical territoryis the small pool of potential hires. “The really strong people don't want to go into a closed environment where it's all secret,” Bengio says. “The differentiating factors are, 'Who are you going to be working with?' 'Am I going to stay a part of the scientific community?' 'How much freedom will I have?' ”

Besides alienating the industry's stars, Apple's secrecy risks turning off promising graduate students, says Trevor Darrell, managing director of a machine-learning research center at the University of California at Berkeley. The ability to continue publishing and otherwise maintain a presence in the scientific community is the most important factor for top students making career decisions, he says. Says Sergey Levine, a research scientist at Google and a postdoctoral researcher at Berkeley: “It's very hard to do science like that.” On Oct. 22, Google announced what it's calling a residency program focused on AI research and publication to further tempt experts.

Apple is slowly expanding its AI staff, acquiring startups such as Perceptio and VocalIQ, as well as hiring influential researchers from companies including Microsoft. Its \href{https://jobs.apple.com/us/search?#&ss=Artificial\%20Intelligence&t=0&so=&lo=0\%2AUSA&pN=0}{jobs website} lists 42 open positions that mention artificial intelligence and 120 that include the words “machine learning.” And AI researchers say they've heard Apple is planning to publish its first major AI paper, but they couldn't provide details.

        In September, Amazon.com, one of the few corporations that \href{http://www.bloomberg.com/bw/articles/2014-06-17/inside-the-secretive-r-and-d-lab-behind-the-amazon-phone}{rivals} Apple's extreme secrecy,
let one of its researchers publish a paper proposing a more efficient way to build AI systems that can recognize images and speech. (An internal committee had to clear everything that made it into the paper, says a person familiar with the matter. Amazon declined to comment.) Ultimately, any company that wants to make its software smarter will have to accept competitors learning some of its tricks, says Matt Zeiler, chief executive officer of AI startup Clarifai. Most researchers in the field, he says, “really want to talk about what they're working on.”

 \emph{\textbf{The bottom line:} Apple is ramping up AI efforts, but the company's reticence to publish its research is limiting its effectiveness and hiring.}`,
			Want: `
\begin{Figure}
        \centering
        \includegraphics[width=0.95\textwidth]{/home/frank/rol/615016607.jpeg}
        \captionof{figure}{Photographer: David Paul Morris/Bloomberg \href{https://assets.bwbx.io/images/users/iqjWHBFdfxIU/iCbVQnix8bsQ/v1/640x-1.jpg}{[Image Source]}}
        \label{fig:615016607}
\end{Figure}

        \href{https://subscribe.businessweek.com/servlet/OrdersGateway?cds_mag_code=BWK&cds_page_id=205566}{}

In the world of artificial intelligence, one of the year's biggest coming-out parties is the Neural Information Processing Systems conference. Thousands of researchers from universities and software companies gather to share their work and wrestle with new ways to tailor software to people's habits. At last year's conference in Montreal, employees of Google, Microsoft, and IBM presented papers on teaching computers to work faster and smarter, such as by reading the house numbers in a photo to determine an address. But one player was conspicuously absent: Apple. This year, Chinese search giant Baidu and Facebook, along with Google and Microsoft, are slated to present papers. Apple isn't.

Apple researchers attended the Montreal conference last year but kept a low profile and didn't say who they worked for unless asked, says Yoshua Bengio, an AI pioneer and professor of computer science at the University of Montreal. This is typical of the company's appearances at the field's big AI conferences, say Bengio and other prominent researchers. “Apple is off the scale in terms of secrecy,” says Richard Zemel, a
professor in the computer science department at the University of Toronto. “They're completely out of the loop.” Apple declined to comment for this story.

Other big consumer-software companies have set up \href{http://www.bloomberg.com/news/features/2015-06-09/the-future-of-computers-is-the-mind-of-a-toddler}{research centers} staffed with dozens or hundreds of AI
experts from around the world, racing to publish findings. The fast-paced culture has produced Google algorithms that can more accurately transcribe speech and tag photos; personal assistants that are smarter than Siri, such as Microsoft's \href{http://www.bloomberg.com/news/videos/b/c2986a39-ede9-4907-8226-fa28a6a57086}{Cortana} and Google's \href{http://www.bloomberg.com/news/articles/2015-05-28/what-google-just-announced-is-a-bombshell}{Now}; and Facebook software that can tell blind users what's in their friends' photos.
\begin{quotation}

 \textbf{0}

Number of AI papers Apple researchers have published
\end{quotation}

At Apple, new hires on the AI teams are told not to announce their positions on LinkedIn or Twitter, says Graham Taylor, a professor of machine learning at Ontario's University of Guelph. They've been told to lock their office doors whenever they leave, says a person familiar with the matter. The teams working to make the company's software smarter are kept unaware of what similar Apple teams are doing, say two people who've worked with the company.

So far, \href{http://www.bloomberg.com/news/videos/2014-09-25/apple-iphone-bends-crashes-into-software-road-bump}{so Apple}. For years, the fortunes of the world's most valuable company have been associated with
its opacity. But its biggest AI success to date has been buying Siri from a \href{http://www.bloomberg.com/bw/articles/2013-02-14/siri-creator-sri-spins-off-a-new-virtual-assistant}{startup} in 2010. Apple Maps still lags the predictive capabilities of similar software. “There's no way they can just observe and not be part of the community and take advantage of what is going on,” Bengio says. “I believe if they don't change their attitude, they will stay behind.”

What makes developing AI different from a mobile operating systemapart from the uncharted technical territoryis the small pool of potential hires. “The really strong people don't want to go into a closed environment where it's all secret,” Bengio says. “The differentiating factors are, 'Who are you going to be working with?' 'Am I going to stay a part of the scientific community?' 'How much freedom will I have?' ”

Besides alienating the industry's stars, Apple's secrecy risks turning off promising graduate students, says Trevor Darrell, managing director of a machine-learning research center at the University of California at Berkeley. The ability to continue publishing and otherwise maintain a presence in the scientific community is the most important factor for top students making career decisions, he says. Says Sergey Levine, a research scientist at Google and a postdoctoral researcher at Berkeley: “It's very hard to do science like that.” On Oct. 22, Google announced what it's calling a residency program focused on AI research and publication to further tempt experts.

Apple is slowly expanding its AI staff, acquiring startups such as Perceptio and VocalIQ, as well as hiring influential researchers from companies including Microsoft. Its \href{https://jobs.apple.com/us/search?#&ss=Artificial\%20Intelligence&t=0&so=&lo=0\%2AUSA&pN=0}{jobs website} lists 42 open positions that mention artificial intelligence and 120 that include the words “machine learning.” And AI researchers say they've heard Apple is planning to publish its first major AI paper, but they couldn't provide details.

        In September, Amazon.com, one of the few corporations that \href{http://www.bloomberg.com/bw/articles/2014-06-17/inside-the-secretive-r-and-d-lab-behind-the-amazon-phone}{rivals} Apple's extreme secrecy,
let one of its researchers publish a paper proposing a more efficient way to build AI systems that can recognize images and speech. (An internal committee had to clear everything that made it into the paper, says a person familiar with the matter. Amazon declined to comment.) Ultimately, any company that wants to make its software smarter will have to accept competitors learning some of its tricks, says Matt Zeiler, chief executive officer of AI startup Clarifai. Most researchers in the field, he says, “really want to talk about what they're working on.”

 \emph{\textbf{The bottom line:} Apple is ramping up AI efforts, but the company's reticence to publish its research is limiting its effectiveness and hiring.}`,
		},
		{
			In: `Test
					\textbf{\\
					In Conclusion:
		
					Capitalism as “Social Pathology”} \\
		
		      \textbf{-STRUCTURAL CLASSISM, THE STATE AND WAR-
		          \\
		          \\
		
		         \footnote{
		         Source: “Man's place in the animal world”,  \emph{What is Man? And other Irreverent Essays, Mark Twain, 1896, p.157}}
		          \\
		         - Mark Twain}  \textbf{Overview}
			`,
			Want: `Test
					\textbf{\\
					In Conclusion:
		
					Capitalism as “Social Pathology”} \\
		
		      \textbf{-STRUCTURAL CLASSISM, THE STATE AND WAR-
		          \\
		          \\
		
		         \footnote{
		         Source: “Man's place in the animal world”,  \emph{What is Man? And other Irreverent Essays, Mark Twain, 1896, p.157}}
		          \\
		         - Mark Twain}  \textbf{Overview}
			`,
		},
	}

	for _, test := range tests {
		got := EmptyLinesInMacros(test.In)
		if err := testutils.MustBeEqual(got, test.Want); err != nil {
			t.Errorf("strings don't match")
		}
	}
	testutils.Cleanup()

}

func Test_SpaceAfterOpeningBrackets__haveMultilinedSpaceAfterBrackets_spaceAndNewlinesAreRemoved(t *testing.T) {
	tests := testutils.ConversionTests{
		{
			In: `Test.
			Test. \emph{
				
				Test}Test
			
			Test`,
			Want: `Test.
			Test. \emph{Test}Test
			
			Test`,
		},
	}

	for _, test := range tests {
		got := SpaceAfterOpeningBrackets(test.In)
		if !reflect.DeepEqual(test.Want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall SpaceBeforeClosingBrackets(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.In, test.Want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}
