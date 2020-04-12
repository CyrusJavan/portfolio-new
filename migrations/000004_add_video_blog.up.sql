INSERT INTO article (author_id, slug, title, content, snippet)
VALUES ((SELECT id FROM author WHERE name = 'Cyrus Javan'), 'video-formats-codecs', 'How does video work on the web?',
$Tag$<p>You probably interact with video online <a href="https://www.statista.com/statistics/319688/daily-online-video-usage/">every day</a> but do you really understand how it works? I thought I did. Until recently when I was implementing video messaging into a web app and I quickly realized how little I knew. A simple question, "What format should a video be in so any device can watch it?", brought me down the rabbit hole that is video on the web. In this article I provide a high level overview of video topics. Specifically highlighting some vocabulary that was confusing and priming your brain to go deeper on any of these topics if you want.</p>
<p><strong>We will cover:</strong></p>
<ul class="list-group">
<li class="list-group-item"><a href="#anatomy-of-a-video-codecs-and-container-formats">Anatomy of a video: Codecs and Container Formats</a></li>
<li class="list-group-item"><a href="#streaming-vs-progressive-downloads">Streaming vs Progressive Downloads</a></li>
<li class="list-group-item"><a href="#transcoding">Transcoding</a></li>
<li class="list-group-item"><a href="#browsers">Browsers</a></li>
</ul>
<div class="mt-5">
<h2 id="anatomy-of-a-video-codecs-and-container-formats">Anatomy of a video: Codecs and Container Formats</h2>
</div>
<div class="d-flex justify-content-center mt-3">
<figure class="figure"><img class="figure-img img-fluid rounded" src="../static/img/multimedia_container.svg" alt="Multimedia container diagram" />
<figcaption class="figure-caption">The layout of a multimedia container.</figcaption>
</figure>
</div>
<p>If a video is an MP4 that means the container format is MP4. The <strong>container format</strong> decides how the data inside the file is organized. The container format does not indicate how the actual audio or video data is encoded or compressed. Examples of container formats are <a href="https://en.wikipedia.org/wiki/WebM">WebM</a>, <a href="https://en.wikipedia.org/wiki/MPEG-4_Part_14">MP4</a> and <a href="https://en.wikipedia.org/wiki/Matroska">Matroska</a>.</p>
<div class="alert alert-info"><strong>History:</strong> One of the very first multimedia file formats was the <a href="https://en.wikipedia.org/wiki/Interchange_File_Format">Interchange File Format (IFF)</a> developed by Electronic Arts in 1985. The format&rsquo;s design was partly inspired by the format Apple&rsquo;s Macintosh&rsquo;s were using for their clipboard. You can check out the <a href="http://www.martinreddy.net/gfx/2d/IFF.txt">original IFF spec</a> which is actually a pretty interesting read as far as technical documents go.</div>
<p>There are 3 things inside the container: metadata, video data and audio data. Metadata tells us a lot about what is going on in the container. Here is the output from <code>mediainfo test.mkv</code> for a video on my computer:</p>
<div>
<pre class="pre-scrollable ws-preline border border-dark rounded px-1"><code>General
Complete name                            : test.mkv
Format                                   : Matroska
Format version                           : Version 4
File size                                : 792 KiB
Writing application                      : Chrome
Writing library                          : Chrome
IsTruncated                              : Yes
FileExtension_Invalid                    : mkv mk3d mka mks

Video
ID                                       : 2
Format                                   : AVC
Format/Info                              : Advanced Video Codec
Codec ID                                 : V_MPEG4/ISO/AVC
Width                                    : 640 pixels
Height                                   : 480 pixels
Display aspect ratio                     : 4:3
Frame rate mode                          : Variable
Language                                 : English
Default                                  : Yes
Forced                                   : No

Audio
ID                                       : 1
Format                                   : Opus
Codec ID                                 : A_OPUS
Channel(s)                               : 1 channel
Channel layout                           : C
Sampling rate                            : 48.0 kHz
Bit depth                                : 32 bits
Compression mode                         : Lossy
Delay relative to video                  : 59 ms
Language                                 : English
Default                                  : Yes
Forced                                   : No
</code></pre>
</div>
<p>We can see that the container format is Matroska, and the video data is in Advanced Video Coding (AVC) format and the audio data is in Opus format. These video and audio formats are known as <strong>codecs</strong>. The codec (an amalgam of the words <strong>co</strong>der and <strong>dec</strong>oder) is the algorithm that is used to encode and decode the media data. Examples of audio codecs are <a href="https://en.wikipedia.org/wiki/Advanced_Audio_Coding">AAC</a> and <a href="https://en.wikipedia.org/wiki/Opus_(audio_format)">Opus</a>. Examples of video codecs are <a href="https://en.wikipedia.org/wiki/Advanced_Video_Coding">AVC/H.264</a>, <a href="https://en.wikipedia.org/wiki/High_Efficiency_Video_Coding">HEVC/H.265</a> and <a href="https://en.wikipedia.org/wiki/VP9">VP9</a>. There are <strong>many</strong> other codecs out there, however, unless you are doing very specific codec work (like trying to improve Netflix&rsquo;s encoding) then you can just stick with the widely used and supported ones.</p>
<p>I will not attempt to describe the details of any codecs here as that is very far out of my wheelhouse, but the main things to understand are:</p>
<ul>
<li><a href="https://en.wikipedia.org/wiki/Comparison_of_video_container_formats">Different container formats can hold different codecs</a></li>
<li>Browsers can only play a subset of all codecs and formats</li>
</ul>
<div class="mt-5">
<h2 id="streaming-vs-progressive-downloads">Streaming vs Progressive Downloads</h2>
</div>
<p>When we use a simple container format like MP4 and point a video element at it the browser will begin a <strong>progressive download</strong> this means the browser will start downloading the video into memory from start to finish. If a user tries to seek to different spot in the video the browser will request that part of the file from the server and continue downloading from there. This method is memory intensive on the clients machine because the browser will attempt to hold the entire video in memory. This makes progressive download unsuitable for longer videos. To avoid buffering the video in memory what you want is a <strong>stream</strong>.</p>
<p>A streaming protocol, like <a href="https://en.wikipedia.org/wiki/HTTP_Live_Streaming">HTTP Live Streaming (HLS)</a>, utilizes the same containers and codecs that you would find in a regular video file, but it will chop the data into bite size chunks. So instead of a single file your video is represented as a directory with a manifest file and the chunks of data. To play a stream the browser reads the manifest file to find the locations of the chunks, then begins requesting the data. The browser will play the data as soon as it is received and does not keep already played chunks in memory. Therfore, the memory impact on the client is the same for a long video as a short video.</p>
<p>An optimization that streaming protocols support is <a href="https://en.wikipedia.org/wiki/Adaptive_bitrate_streaming">Adaptive Bitrate Streaming (ABR)</a>. With ABR we create multiple different versions of those data chunks, each encoded at a different <a href="https://filmora.wondershare.com/video-editing-tips/what-is-video-bitrate.html">bitrate</a> (lower bitrate means lower quality). The browser will request data chunks at the highest bitrate it's internet connection can handle without having choppy video. If the browser is experiencing choppy video, it will request data at a lower bitrate to smooth out the video.</p>
<div class="mt-5">
<h2 id="transcoding">Transcoding</h2>
</div>
<p>The process of taking a video from one format to another is known as <strong>transcoding</strong>. If you want to convert a user uploaded WebM video into an adaptive bitrate HLS stream you will need to transcode. Transcoding works by decoding the video to a raw (uncompressed) format then encoding in the desired format. The principle of <a href="https://en.wikipedia.org/wiki/Garbage_in,_garbage_out">"garbage in, garbage out"</a> applies here, you can never transcode to a higher quality than what you started with. The standard tool for transcoding is <a href="https://www.ffmpeg.org/">ffmpeg</a><span style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;"> which has a huge amount of options and can run pretty much anywhere. However, if you have a large amount of videos maybe you don&rsquo;t want to deal with having to run ffmpeg as a service you instead could use a hosted 3rd party solution. AWS offers their </span><a style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;" href="https://aws.amazon.com/mediaconvert/">MediaConvert</a><span style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;"> service which hooks in nicely with S3 and CloudFront. There are also companies who solely do transcoding like </span><a style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;" href="https://zencoder.com/en/file-transcoding/pricing">Zencoder</a><span style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;"> which could also be a good option.</span></p>
<div class="mt-5">
<h2 id="browsers">Browsers</h2>
</div>
<p>Browsers are picky about what streaming protocols, container formats and codecs they are willing to work with. This is really where you need to pay attention as a web developer. A useful resource is the Mozilla Developer Network (MDN) <a href="https://developer.mozilla.org/en-US/docs/Web/Media/Formats">media type and format guide</a>, which has information on browser support.</p>
<div class="d-flex justify-content-center mt-3">
<figure class="figure"><img class="figure-img img-fluid rounded" style="max-width: 400px;" src="../static/img/no_supported_format.png" alt="No supported format from firefox" />
<figcaption class="figure-caption">This looks pretty bad to your users.</figcaption>
</figure>
</div>
<div class="alert alert-info"><strong>Note:</strong> Browsers implement the function <a href="https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/canPlayType">canPlayType</a> which takes one parameter, a string of a MIME type, and returns a string response which tells you if it can play the video. Due to the diverse nature of container formats and codecs, the browser will only give one of three responses: "" (empty string, meaning no the browser can't play the video), "maybe" and "probably".</div>
<p>To answer the original question,&nbsp;<span style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;">"What format should a video be in so any device can watch it?", </span><span style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;">the best answer we can give is an MP4 container format with the codecs Advanced Audio Coding (AAC) for audio and AVC/H.264 for video. For streaming, the HLS protocol is supported by every major browser.</span></p>
<p>If you want to go deeper into the rabbit role of formats and codecs, many of their specifications are open online. To learn more about best practices around using video on the web <a href="https://developers.google.com/web/fundamentals/media/video">this article</a> from Google's Web Fundamentals series is great. Hopefully this article gave you a better understanding of the basics of video on the web, thanks for reading.</p>$Tag$,
$$You probably interact with video online every day but do you really understand how it works? I thought I did. Until recently when I was implementing video messaging into a web app and I quickly realized how little I knew. A simple question, "What format should a video be in so any device can watch it?", brought me down the rabbit hole that is video on the web. In this article I provide a high leve...$$);

INSERT INTO tag (name) VALUES ('video');
INSERT INTO tag (name) VALUES ('web development');

INSERT INTO article_tag (article_id, tag_id)
VALUES ((SELECT id FROM article WHERE slug = 'video-formats-codecs'),
(SELECT id FROM tag WHERE name = 'video'));

INSERT INTO article_tag (article_id, tag_id)
VALUES ((SELECT id FROM article WHERE slug = 'video-formats-codecs'),
(SELECT id FROM tag WHERE name = 'web development'));