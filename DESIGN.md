{\rtf1\ansi\ansicpg1252\cocoartf1504\cocoasubrtf830
{\fonttbl\f0\fswiss\fcharset0 Helvetica;\f1\froman\fcharset0 TimesNewRomanPSMT;\f2\fnil\fcharset0 HelveticaNeue;
}
{\colortbl;\red255\green255\blue255;\red53\green53\blue53;\red191\green191\blue191;}
{\*\expandedcolortbl;;\csgenericrgb\c20784\c20784\c20784;\csgray\c79525;}
\margl1440\margr1440\vieww10800\viewh8400\viewkind0
\deftab560
\pard\pardeftab560\ri0\qj\partightenfactor0

\f0\fs24 \cf0 When you enter the site, you will first be presented with an event category video to attract users' attention. Click \'93Let's begin\'94 to enter the activity display page.\
\
On the front page "What are you looking for? "Section, which presents the user with five different event types, and the user can click on it to view all the events under that type. In the "Check out what's happening" section of the homepage, 6 different activities will be displayed randomly, and the activities presented here will change every time the page is refreshed. The "All Events" section at the bottom of the home page shows all the activities. For each activity displayed on the homepage, users can click to enter the activity details page to obtain more detailed activity information.\
\
There is a search box in the upper right corner of the home page that allows users to search for activities they are interested in. After entering the search results page, users can find the filter box at the top left of the page. Users can filter the search results by event type and event time.\
\
There is also a create event button in the upper right corner of the home page, which users can click to create new activities. In creating the activity page, the user enters the activity name, location, picture, date, category, and activity details to create the page. At the same time, the page will automatically check the user's input. Valid titles must be more than 5 and fewer than 50 unicode characters. A valid location must be More than 5 and fewer than 50 unicode characters. Valid images should be in a URL and be one of the following file types-".png", ".jpg", ". GIF\'94, \'93.jpeg", or ". Gifv ". Finally, valid Dates must be in the format "2006-01-02 t15:04 "And also be in the future. If any of these requirements are incorrect, the page will return the user to the form (at the same URL) and show them their error.\
\
In the event details page, users can learn about the event, location and other details of the event. The user can also enter an email address to book an event, and the page only allows you to register with @yale.edu. If the users enter an invalid email, the page will show them it\'92s an invalid email. Once the users enters a valid email,  they will be given a confirmation code which is the first seven characters of the sha256 hash of their email address. \
\
And the user can click "Want to support?" to make a donation, click the button and a heart-shaped animation will appear to thank the user for the donation.\
\
Our website is also mobile friendly, it can adapt to mobile browsers.
\f1 \
\pard\pardeftab560\ri0\partightenfactor0
\cf0 \
\pard\pardeftab560\ri0\partightenfactor0

\f0 \cf0 We design and create web pages using HTML, CSS, JavaScript, and the GO language. Here's an overview of what each file does.
\f1 \
\pard\pardeftab560\ri0\partightenfactor0

\f2 \cf2 \

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrt\brdrs\brdrw10\brdrcf3 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./server.go\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 Entrypoint for the code; contains the\'a0main\'a0function\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./routes.go\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 Maps our URLs to the controllers/handlers for each URL\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./event_models.go\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 Defines our data structure and logic\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./index_controllers.go	\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 Controllers related to the index (home) page\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./templates.go\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 Sets up our templates from which you generate HTML\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./static.go\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 Sets up the static file server (see next entry)\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./staticfiles\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 Directory with our "static" assets like images, css, js\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./templates/layout.gohtml\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 The "base" layout for your HTML frontend\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./templates/index.gohtml\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 The template for the index (home) page\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./templates/event-detail.gohtml\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 The template for the event detail page (when viewing a single event)\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./templates/event-new.gohtml\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 The layout when a user creates a new event\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./templates/search.gohtml\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 The layout when a user searches for existing events\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./templates/searchcategory.gohtml\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 The layout when the user selects a category from the homepage and sees the list of all available events under a specific category\cell \row

\itap1\trowd \taflags1 \trgaph108\trleft-108 \trbrdrl\brdrs\brdrw10\brdrcf3 \trbrdrb\brdrs\brdrw10\brdrcf3 \trbrdrr\brdrs\brdrw10\brdrcf3 
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx4320
\clvertalt \clshdrawnil \clwWidth4675\clftsWidth3 \clbrdrt\brdrs\brdrw10\brdrcf3 \clbrdrl\brdrs\brdrw10\brdrcf3 \clbrdrb\brdrs\brdrw10\brdrcf3 \clbrdrr\brdrs\brdrw10\brdrcf3 \clpadl100 \clpadr100 \gaph\cellx8640
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 ./templates/aboutpage.gohtml\cell 
\pard\intbl\itap1\pardeftab560\ri0\partightenfactor0
\cf2 About page for more info about our team\cell \lastrow\row
\pard\pardeftab560\ri0\partightenfactor0
\cf2 \
}