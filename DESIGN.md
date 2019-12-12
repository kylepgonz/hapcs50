Project Name: Hap
Members: Kyle Gonzales, Lucas Qiang Wei

When you enter the site, you will first be presented with an event category video to attract users' attention. Click “Let's begin” to enter the activity display page.

On the front page "What are you looking for? "Section, which presents the user with five different event types, and the user can click on it to view all the events under that type. In the "Check out what's happening" section of the homepage, 6 different activities will be displayed randomly, and the activities presented here will change every time the page is refreshed. The "All Events" section at the bottom of the home page shows all the activities. For each activity displayed on the homepage, users can click to enter the activity details page to obtain more detailed activity information.

There is a search box in the upper right corner of the home page that allows users to search for activities they are interested in. After entering the search results page, users can find the filter box at the top left of the page. Users can filter the search results by event type and event time.

There is also a create event button in the upper right corner of the home page, which users can click to create new activities. In creating the activity page, the user enters the activity name, location, picture, date, category, and activity details to create the page. At the same time, the page will automatically check the user's input. Valid titles must be more than 5 and fewer than 50 unicode characters. A valid location must be More than 5 and fewer than 50 unicode characters. Valid images should be in a URL and be one of the following file types-".png", ".jpg", ".GIF", ".jpeg", or ".Gifv". Finally, valid Dates must be in the format "2006-01-02 t15:04 "And also be in the future. If any of these requirements are incorrect, the page will return the user to the form (at the same URL) and show them their error.

In the event details page, users can learn about the event, location and other details of the event. The user can also enter an email address to book an event, and the page only allows you to register with @yale.edu. If the users enter an invalid email, the page will show them it’s an invalid email. Once the users enters a valid email,  they will be given a confirmation code which is the first seven characters of the sha256 hash of their email address. 

And the user can click "Want to support?" to make a donation, click the button and a heart-shaped animation will appear to thank the user for the donation.

Our website is also mobile friendly, it can adapt to mobile browsers.

We design and create web pages using HTML, CSS, JavaScript, and the GO language. Here's an overview of what each file does.

| File                              | Description                                                                                                                       |
|-----------------------------------|-----------------------------------------------------------------------------------------------------------------------------------|
| ./server.go                       | Entrypoint for the code; contains the main function                                                                               |
| ./routes.go                       | Maps our URLs to the controllers/handlers for each URL                                                                            |
| ./event_models.go                 | Defines our data structure and logic                                                                                              |
| ./index_controllers.go            | Controllers related to the index (home) page                                                                                      |
| ./templates.go                    | Sets up our templates from which you generate HTML                                                                                |
| ./static.go                       | Sets up the static file server (see next entry)                                                                                   |
| ./staticfiles                     | Directory with our "static" assets like images, css, js                                                                           |
| ./templates/layout.gohtml         | The "base" layout for your HTML frontend                                                                                          |
| ./templates/index.gohtml          | The template for the index (home) page                                                                                            |
| ./templates/index.gohtml          | The template for the event detail page (when viewing a single event)                                                              |
| ./templates/event-new.gohtml      | The layout when a user creates a new event                                                                                        |
| ./templates/search.gohtml         | The layout when a user searches for existing events                                                                               |
| ./templates/searchcategory.gohtml | The layout when the user selects a category from the homepage and sees the list of all available events under a specific category |
| ./templates/aboutpage.gohtml      | About page for more info about our team                                                                                           |
