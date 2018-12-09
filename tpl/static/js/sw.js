/**
 * Welcome to your Workbox-powered service worker!
 *
 * You'll need to register this file in your web app and you should
 * disable HTTP caching for this file too.
 * See https://goo.gl/nhQhGp
 *
 * The rest of the code is auto-generated. Please don't update this file
 * directly; instead, make changes to your Workbox build configuration
 * and re-run your build process.
 * See https://goo.gl/2aRDsh
 */

importScripts("https://storage.googleapis.com/workbox-cdn/releases/3.6.3/workbox-sw.js");

/**
 * The workboxSW.precacheAndRoute() method efficiently caches and responds to
 * requests for URLs in the manifest.
 * See https://goo.gl/S9QRab
 */
self.__precacheManifest = [
  {
    "url": "2018/04/12/beejay-bakau-resort.html",
    "revision": "e725f093063d0f6c884f47f7ae946c1e"
  },
  {
    "url": "2018/05/27/sewa-mobil-di-surabaya.html",
    "revision": "eeec618164703c8db3a53a23fb9fd136"
  },
  {
    "url": "2018/06/21/harga-sewa-mobil-surabaya.html",
    "revision": "6c596c93b13eb41b83551712734f9abf"
  },
  {
    "url": "2018/07/23/sewa-mobil-bandara-juanda.html",
    "revision": "131981acd5a81b865f179a697a1f95b3"
  },
  {
    "url": "2018/08/16/sewa-mobil-tanggulangin.html",
    "revision": "840200fb78bd0db20315e3579f420774"
  },
  {
    "url": "2018/09/28/sewa-mobil-elf-surabaya.html",
    "revision": "968870bf7509e72613383f5c916be97a"
  },
  {
    "url": "2018/10/05/rental-annual-meeting-imf-world-bank-di-bali.html",
    "revision": "c26b88461e7c52bf20d9883b5f6d205c"
  },
  {
    "url": "2018/11/05/keunggulan-rental-mobil-surabaya.html",
    "revision": "7fd5cd2b24b871c4b9b446a265e92c6a"
  },
  {
    "url": "404.html",
    "revision": "9bc6b56e88076cbc0999e59f58115d31"
  },
  {
    "url": "blog.html",
    "revision": "057db4369e71a33f7f3b8e3f42c35e56"
  },
  {
    "url": "contact.html",
    "revision": "1363f8d2bce9f14167fa0140801e508b"
  },
  {
    "url": "favicon.ico",
    "revision": "b40651e0f537b50a715caec3caaeb015"
  },
  {
    "url": "feed.xml",
    "revision": "0ac08a82613deafc630a8beded1ebfa5"
  },
  {
    "url": "gallery.html",
    "revision": "5faa98cfd49118b36bc2960dd88acb6a"
  },
  {
    "url": "google80eb15f70004671a.html",
    "revision": "2c06e59ed42caa2f9c628716dd34fc23"
  },
  {
    "url": "images/icons/icon-128x128.png",
    "revision": "6fff73ab8a11fd520523ebac296822f2"
  },
  {
    "url": "images/icons/icon-144x144.png",
    "revision": "8e372daf020abd67fb6246c825947955"
  },
  {
    "url": "images/icons/icon-152x152.png",
    "revision": "a13cb3800016ce80dcb62ed9c3f1e012"
  },
  {
    "url": "images/icons/icon-192x192.png",
    "revision": "93846d18b02afc5edd56d0b7cca1d12e"
  },
  {
    "url": "images/icons/icon-384x384.png",
    "revision": "e50bffea785bbb1e74540955e7d4105c"
  },
  {
    "url": "images/icons/icon-512x512.png",
    "revision": "8a1162e0405dc783d0197e787dc7be2d"
  },
  {
    "url": "images/icons/icon-72x72.png",
    "revision": "0f816e4baa8f03a54948c6485e8f8fb0"
  },
  {
    "url": "images/icons/icon-96x96.png",
    "revision": "5a56f0ca5394877f21ed88e607662fec"
  },
  {
    "url": "index.html",
    "revision": "42a6829e658ed111373b6f76ad0e9687"
  },
  {
    "url": "manifest.json",
    "revision": "5c4bcd2f65b1b91cf9cca73c1bb17843"
  },
  {
    "url": "photos/imf-1.jpg",
    "revision": "3ce07264e73f942a45df051ceb830be9"
  },
  {
    "url": "photos/imf-2.jpg",
    "revision": "b1a422c3e6818fab43e50f5401e101a3"
  },
  {
    "url": "photos/november-1.jpg",
    "revision": "af6579fcc96347100a8409a8fcf5d2d6"
  },
  {
    "url": "product.html",
    "revision": "740273fe38e4906abe2330aa644bc29a"
  },
  {
    "url": "/s/robots.txt",
    "revision": "1a2941002cfcb13466576b3dcdbfb96f"
  },
  {
    "url": "sitemap.xml",
    "revision": "2021a3bf137781fd9b144ca9661d9436"
  },
  {
    "url": "static/32385489_1021643167993325_7243333418375708672_n.jpg",
    "revision": "586820ef5f903c94209d2d8b3e670fd8"
  },
  {
    "url": "static/41630686_345660996007408_3959362786232315390_n.jpg",
    "revision": "8fadc4f1ca68d70da0c6d06d008dcb7c"
  },
  {
    "url": "static/41816743_335570487006518_1070829132906493154_n.jpg",
    "revision": "15a469ba1fa7aae3aa80d422de623c92"
  },
  {
    "url": "static/42003704_2169489999956288_4874548311978663431_n.jpg",
    "revision": "8bda534addddcfb26a5a0596cbd0fbae"
  },
  {
    "url": "static/43061223_323731011736020_6734971322820834699_n.jpg",
    "revision": "6300e4b7084402cfe83128af370129cf"
  },
  {
    "url": "static/43597816_333862287374890_4164960330071885602_n.jpg",
    "revision": "7b1706e70abe5e31037aba2dccb7641f"
  },
  {
    "url": "static/43767628_2182134858675310_4593096987499668786_n.jpg",
    "revision": "f9fd611078305b722b90c89ec807d2e0"
  },
  {
    "url": "static/43817648_353828892048143_45810111997129597_n.jpg",
    "revision": "a4dc1c1360fd55868df899f323e67127"
  },
  {
    "url": "static/43913352_268369640351497_6468748527204336937_n.jpg",
    "revision": "000e7ff916df36f917bd8ce927db2269"
  },
  {
    "url": "static/address.png",
    "revision": "91770a002238b7a78f861d022faa084c"
  },
  {
    "url": "static/alphard.jpg",
    "revision": "c119c0a23ecb43431d1b1bd5d8834af3"
  },
  {
    "url": "static/Audiowide-Regular.ttf",
    "revision": "fb04830aee1eff6efdfeca9c1dca0304"
  },
  {
    "url": "static/avanza.jpg",
    "revision": "b940c14a90b8860a6182857ff8fc4c82"
  },
  {
    "url": "static/booking.png",
    "revision": "1629ac62627e0d7100746972eb1a29c1"
  },
  {
    "url": "static/call.png",
    "revision": "0e78a526f471441c48286cd08495dee4"
  },
  {
    "url": "static/camry.jpg",
    "revision": "737d86636c2f586a08ed3e7e601f28ee"
  },
  {
    "url": "static/cheap.png",
    "revision": "3c551722223f5a64d7838eec50540ad9"
  },
  {
    "url": "static/close.svg",
    "revision": "3360c57cdbbebca82e3a792f0a57876c"
  },
  {
    "url": "static/contact.png",
    "revision": "8044224ec57df42fa3fc1e5b287b9e01"
  },
  {
    "url": "static/elf-long.jpg",
    "revision": "dd591d159e371ba15a8fb4ddb3fe47d8"
  },
  {
    "url": "static/facebook.png",
    "revision": "97eb2fbd6abdc382e86b99b90bab6b6e"
  },
  {
    "url": "static/favicon.ico",
    "revision": "b40651e0f537b50a715caec3caaeb015"
  },
  {
    "url": "static/fortuner-2.jpg",
    "revision": "200a26f1f73fa56b7a92052dfbff4f6c"
  },
  {
    "url": "static/fortuner.jpg",
    "revision": "ce9b71d674dbcd92f5c52e8315a2c87b"
  },
  {
    "url": "static/friend.png",
    "revision": "5145708211d1de9ae5652d9abac36850"
  },
  {
    "url": "static/gallery-photo.png",
    "revision": "b10377e0f7078637b2e408c87bb1a41a"
  },
  {
    "url": "static/google-plus.png",
    "revision": "fc541cefc9fd5a6833d8e0d78e301f23"
  },
  {
    "url": "static/grand-livina.jpg",
    "revision": "ce35ece3d498cb3624e4f0949b574073"
  },
  {
    "url": "static/hamburger.svg",
    "revision": "c28b192512aeeef7107d50fc1d61dc85"
  },
  {
    "url": "static/hiace.jpg",
    "revision": "2be6370b8b58d883004a8bd331156d45"
  },
  {
    "url": "static/innova.jpg",
    "revision": "bd371f397e39fbf22ec79bc2b0f64b10"
  },
  {
    "url": "static/instagram.png",
    "revision": "1b9bf1b1307a06fae046eec3d7e1c85f"
  },
  {
    "url": "static/jemput.png",
    "revision": "b1c14924374fe1eb6d8de9028b9724b0"
  },
  {
    "url": "static/man.png",
    "revision": "9010188083d12b533adee54364d8464f"
  },
  {
    "url": "static/manbrown.png",
    "revision": "f216ea7bc366282a7e655e9d06545a3a"
  },
  {
    "url": "static/news-paper.png",
    "revision": "08f07a9256fc0530039d05c5274a38e2"
  },
  {
    "url": "static/notes.png",
    "revision": "9c2d061249f9eee8e757fc20a890e886"
  },
  {
    "url": "static/phone-book.png",
    "revision": "46d9c55bd7a5146ab49acb1e07f102e0"
  },
  {
    "url": "static/phone.png",
    "revision": "c6e5a64e867b87b8092819b12d16666b"
  },
  {
    "url": "static/phones-symbol.svg",
    "revision": "f16134d24c95484a04754e56c6b0986c"
  },
  {
    "url": "static/rental-mobil-surabaya-1-1.jpg",
    "revision": "c1a702d1d282cc21b426a9b704bbd6b3"
  },
  {
    "url": "static/rental-mobil-surabaya-1-2.jpg",
    "revision": "9dd83ee3ec9b31f65aaa7afb2783f227"
  },
  {
    "url": "static/rental-mobil-surabaya-1.jpg",
    "revision": "bc91238fed583e33ab0488047732331c"
  },
  {
    "url": "static/rental-mobil-surabaya-2.jpg",
    "revision": "e9ab4812d3bb2fcd8bd9ccd8cd05098c"
  },
  {
    "url": "static/rental-mobil-surabaya-3-1.jpg",
    "revision": "2b5b244a9a731a9c88b9cfe6812c87f8"
  },
  {
    "url": "static/rental-mobil-surabaya-3-2.jpg",
    "revision": "08423598d6680832c0efe9e2c2f724d0"
  },
  {
    "url": "static/rental-mobil-surabaya-3.jpg",
    "revision": "be2ac9f4d451390d9aa1131aada05351"
  },
  {
    "url": "static/rental-mobil-surabaya-4.jpg",
    "revision": "11c4fed67e46b9797b103e43cbb26b10"
  },
  {
    "url": "static/rental-mobil-surabaya-tran99-logo.png",
    "revision": "2a5d252607f2dd50e5f310d5289d94d5"
  },
  {
    "url": "static/rental-mobil-surabaya-transparan.jpg",
    "revision": "8af97082a72ac69ff9df359e3c38bb29"
  },
  {
    "url": "static/rental-mobil-surabaya-transparan.png",
    "revision": "326d06ba2610624a43d1308182a87425"
  },
  {
    "url": "static/rental-mobil-surabaya.jpg",
    "revision": "6acbb9558d35644f0471931b80cd8ad5"
  },
  {
    "url": "static/rental-mobil-surabaya.png",
    "revision": "3f1c2887e1f13cf5eabbd940d298cb94"
  },
  {
    "url": "static/Roboto-Regular.ttf",
    "revision": "3e1af3ef546b9e6ecef9f3ba197bf7d2"
  },
  {
    "url": "static/sms.svg",
    "revision": "878cb683823bc63731f8ef65323abc48"
  },
  {
    "url": "static/time.png",
    "revision": "83b0ca8a7da7c9af1c40c898d7f75610"
  },
  {
    "url": "static/transfer.png",
    "revision": "895ab9052d59b494501eecc538075af9"
  },
  {
    "url": "static/transfer.svg",
    "revision": "e7f36d62cf818adf32897f65cf5b2788"
  },
  {
    "url": "static/trust.png",
    "revision": "3508a44dea145fa061cc767b2fbf92cb"
  },
  {
    "url": "static/twitter.png",
    "revision": "d9db4e79ec771f7dc9c2eb3198bf8aba"
  },
  {
    "url": "static/whatsapp-2.png",
    "revision": "ddc83346a99e4d3e218d941c67309c07"
  },
  {
    "url": "static/whatsapp.png",
    "revision": "63829b6143f17134ec059387d805ac4a"
  },
  {
    "url": "static/whatsapp.svg",
    "revision": "47c77bf133116ff76838809ae5d2c83a"
  },
  {
    "url": "sw.html",
    "revision": "9ddb3c30045037e3f8d7c4c5810a4386"
  }
].concat(self.__precacheManifest || []);
workbox.precaching.suppressWarnings();
workbox.precaching.precacheAndRoute(self.__precacheManifest, {});
