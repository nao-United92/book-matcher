<h1 align="center"> Book Matcher </h1>
<h5 align="center"> Development Started - Summer 2023 </h5>

[English](https://github.com/nao-United92/book-matcher/blob/main/README.md) | 日本語</h2>

<!-- TABLE OF CONTENTS -->
<h2 id="table-of-contents"> :book: 目次</h2>

<details open="open">
  <summary>目次</summary>
  <ol>
    <li><a href="#overview"> ➤ 概要</a></li>
    <li><a href="#use-stack"> ➤ 使用技術</a></li>
    <li><a href="#project-files-description"> ➤ ファイル概要</a></li>
    <li><a href="#getting-started"> ➤ 起動方法</a></li>
    <li><a href="#ongoing"> ➤ 実装中</a></li>
    <li><a href="#to-be-implemented"> ➤ 実装予定</a></li>
    <li><a href="#references"> ➤ 参考</a></li>
  </ol>
</details>


![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- 概要 -->
<h2 id="overview"> :cloud: 概要</h2>

<p align="justify">
  好きな本で人とつながるコミュニティサービス <b>Book Matcher</b> のプロジェクトリポジトリです。<br>
  <b>※まだまだ未完成です。これからメイン機能の新規実装、既存機能を含む拡張、改修を行なっていきます。</b>
</p>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- 使用技術 -->
<h2 id="use-stack"> :gear: 使用技術</h2>

<p align="justify">
<img src="https://img.shields.io/badge/-Docker-1488C6.svg?logo=docker&style=plastic">
<img src="https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic">
<img src="https://img.shields.io/badge/-Mysql-4479A1.svg?logo=mysql&style=plastic">
<img src="https://img.shields.io/badge/-Html5-E34F26.svg?logo=html5&style=plastic">
<img src="https://img.shields.io/badge/-Css3-1572B6.svg?logo=css3&style=plastic">
<img src="https://img.shields.io/badge/-Jquery-0769AD.svg?logo=jquery&style=plastic">
</p>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- ファイル概要 -->
<h2 id="project-files-description"> :floppy_disk: ファイル概要</h2>

<h3>controller</h3>
<ul>
  <li><b>bookmark/bookmark.go</b> - ブックマーク機能に関連するイベント処理</li>
  <li><b>login/login.go</b> - ログイン機能に関連するイベント処理</li>
  <li><b>password/password.go</b> - パスワード再発行機能に関連するイベント処理</li>
  <li><b>profile/profile.go</b> - プロフィール機能に関連するイベント処理</li>
  <li><b>search/search.go</b> - 書籍検索機能に関連するイベント処理</li>
  <li><b>signup/signup.go</b> - ユーザー新規登録機能に関連するイベント処理</li>
</ul>

<h3>db/my.cnf</h3>
<ul>
  DB(MySQL)サーバの設定ファイルで、サーバ起動時に読み込まれます。DB の動作に関するオプションが設定されています。
</ul>

<h3>model</h3>
<ul>
  <li><b>bookmark.go</b> - ブックマークデータのテーブルカラム定義、データの CRUD 処理</li>
  <li><b>user</b> - ユーザーデータのテーブルカラム定義、データの CRUD 処理</li>
</ul>

<h3>static</h3>
<ul>
  <li><b>css/style.css</b> - 共通 CSS ファイル</li>
  <li><b>images/</b> - テスト用のサンプル画像が格納されています。</li>
  <li><b>js/script.js</b> - 共通 JS ファイル</li>
</ul>

<h3>view</h3>
<ul>
  <li><b>base.html</b> - 共通ページのデザイン</li>
  <li><b>bookmark.html</b> - ブックマークリストフォームのデザイン</li>
  <li><b>login.html</b> - ログインフォームのデザイン</li>
  <li><b>menu.html</b> - メニュー(初期表示)ページのデザイン</li>
  <li><b>password_reset.html</b> - パスワード再発行フォームのデザイン</li>
  <li><b>profile.html</b> - プロフィールフォームのデザイン</li>
  <li><b>search_books.html</b> - 書籍検索リストフォームのデザイン</li>
  <li><b>search_books_detail.html</b> - 書籍詳細フォームのデザイン</li>
  <li><b>signup.html</b> - ユーザー新規登録フォームのデザイン</li>
</ul>

<h3>.env</h3>
<ul>
  データベースの接続情報が記述されています。main.go から読み出され、DB 接続で使用されます。
</ul>

<h3>Dockerfile</h3>
<ul>
  Docker イメージの指定、ホットリロードツール(air)のインストール、WORKDIR の変更を実施
</ul>

<h3>compose.yaml</h3>
<ul>
  アプリサーバ(go)、DB サーバ(db)、DBクライアントツール(phpmyadmin)のコンテナ定義が記述されています。
</ul>

<h3>go.mod</h3>
<ul>
  Go プログラミングで使われるモジュールの依存関係が記述されています。
</ul>

<h3>go.sum</h3>
<ul>
  go.mod ファイル内のモジュールの依存関係をハッシュ化した情報が記述されています。
</ul>

<h3>main.go</h3>
<ul>
  Go プログラムのエントリーポイントです。プログラム起動直後は、まず「main」パッケージに属する「main 関数」が実行されます。
</ul>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- 起動方法 -->
<h2 id="getting-started"> :arrow_forward: 起動方法</h2>

<p><b>1. 初めて起動する場合</b> - イメージ構築、コンテナ構築・起動を一括で実行</p>
<pre><code>$ docker-compose up --build -d</code></pre>

<p><b>2. 起動2回目以降の場合</b> - コンテナ構築・起動を実行</p>
<pre><code>$ docker-compose up -d</code></pre>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- 実装中 -->
<h2 id="ongoing"> :crossed_swords: 実装中</h2>
<ul>
  <li>パスワード再設定機能</li>
</ul>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- 実装予定 -->
<h2 id="to-be-implemented"> :memo: 実装予定</h2>
<ul>
  <li>書籍レビュー投稿機能</li>
  <li>フレンド閲覧・追加機能</li>
  <li>通知機能</li>
</ul>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- 参考 -->
<h2 id="references"> :scroll: 参考</h2>
<a href="https://developers.google.com/books/docs/v1/using?hl=ja">Google Books APIs</a><br>
<a href="https://qiita.com/s_emoto/items/975cc38a3e0de462966a">MVCモデルについて</a><br>
<a href="https://qiita.com/soicchi/items/2637a9195e64fdc73609">go.mod、go.sumファイルは何なのか</a><br>
<a href="https://the-simple.jp/what-is-o-r-mapper-explains-the-basic-concept-of-o-r-mapping-and-how-to-use-it-in-practice">O/Rマッパーとは？O/Rマッピングの基本概念と実践的な使い方を解説</a><br>
<a href="https://fonts.google.com/icons">Material Symbols and Icons - Google Fonts</a>
<br><br>

<a href="https://github.com/nao-United92" target="_blank"><img alt="Github" src="https://img.shields.io/badge/GitHub-%2312100E.svg?&style=for-the-badge&logo=Github&logoColor=white" /></a>
<a href="https://qiita.com/nao-United92" target="_blank"><img alt="Medium" src="https://img.shields.io/badge/qiita-55C500.svg?&style=for-the-badge&logo=qiita&logoColor=white" /></a>
