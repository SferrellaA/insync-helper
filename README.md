# insync-helper
A while ago I got an [Insync](https://www.insynchq.com/) license, but found it lacking in a very particular edge-case. I was using Insync on a Raspberry pi to sync my Google Drive to a network file serve, but found I couldn't open .gdoc files from computers without Insync installed. This was because the .gdoc files created by Insync are just a few lines of text _representing_ the Google Doc, and not the file itself. To let me open .gdoc files from other computers anyway, I made insync-helper that reads the text file and opens the appropriate Google Doc in browser. 

The tool's logic is a bit brittle, and due for a replacement. #todo
