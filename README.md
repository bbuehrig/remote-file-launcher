# remote-file-launcher
Launches files on a mac via a website remotely. I use it for launching songs on my MacBook Pro in Presonus Studio One from my iPad.

You start a file via Webbrowser on your iPad or other devices. Open the site http://<your-ip-adress>:5000 and you get an overview of all files. Start it with a tap on the button.

Configure your files to start via a JSON-file called routes.json:

```
{
    "routes": [{
        "name": "Ant",
        "filepath":  "/S1-Tracks/Ant/Ant.song",
        "urlpath": "ant"
    },  {
        "name": "Don't Look Back",
        "filepath":  "/S1-Tracks/Dont Look Back/Dont Look Back.song",
        "urlpath": "dontlookback"
    },  {
        "name": "Fix",
        "filepath":  "/S1-Tracks/Fix/Fix.song",
        "urlpath": "fix"
    }]
}
```

* __name__: the name of the file for the launching website
* __filepath__: the file that will open on click
* __urlpath__: the GET-Request-ID for this file

After you click on a button on your mobile remote device there will be a GET-request send to the macbook:
```
http://<your-ip-adrees>:5000/open/:urlpath
```

Call it directly to open the file without the overview-page.