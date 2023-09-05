<h1 align="center"> Book Matcher </h1>
<h5 align="center"> Development Started - Summer 2023 </h5>

<!-- TABLE OF CONTENTS -->
<h2 id="目次"> :book: 目次</h2>

<details open="open">
  <summary>目次</summary>
  <ol>
    <li><a href="#overview"> ➤ 概要</a></li>
    <li><a href="#use-stack"> ➤ 使用技術</a></li>
    <li><a href="#project-files-description"> ➤ ファイル概要</a></li>
    <li><a href="#getting-started"> ➤ 起動方法</a></li>
    <li><a href="#function1"> ➤ 機能 1: メニュー </a></li>
    <li><a href="#function2"> ➤ 機能 2: 新規登録 </a></li>
    <li><a href="#function3"> ➤ 機能 3: ログイン </a></li>
    <li><a href="#function4"> ➤ 機能 4: 書籍検索　</a></li>
    <li><a href="#function5"> ➤ 機能 5: プロフィール　</a></li>
    <li><a href="#function6"> ➤ 機能 6: ブックマーク　</a></li>
    <li><a href="#function7"> ➤ 機能 7: ログアウト </a></li>
    <li><a href="#ongoing"> ➤ 実装中</a></li>
    <li><a href="#to-be-implemented"> ➤ 実装予定</a></li>
    <li><a href="#references"> ➤ 参考文献</a></li>
    <li><a href="#link"> ➤ リンク</a></li>
  </ol>
</details>


![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- 概要 -->
<h2 id="overview"> :cloud: 概要</h2>

<p align="justify"> 
  好きな本で人と繋がるマッチングサービス <b>Book Matcher</b> のプロジェクトリポジトリです。
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

<!-- PROJECT FILES DESCRIPTION -->
<h2 id="project-files-description"> :floppy_disk: ファイル概要</h2>

<h3>controller</h3>
<ul>
  <li><b>bookmark/bookmark.go</b> - ブックマーク機能に関連するイベント処理が記述されています。</li>
  <li><b>login/login.go</b> - ログイン機能に関連するイベント処理が記述されています。</li>
  <li><b>password/password.go</b> - パスワード再発行機能に関連するイベント処理が記述されています。</li>
  <li><b>profile/profile.go</b> - プロフィール機能に関連するイベント処理が記述されています。</li>
  <li><b>search/search.go</b> - 書籍検索機能に関連するイベント処理が記述されています。</li>
  <li><b>signup/signup.go</b> - ユーザー新規登録機能に関連するイベント処理が記述されています。</li>
</ul>

<h3>db/my.cnf</h3>
<ul>
  <b>DB(MySQL)サーバの設定ファイルで、サーバ起動時に読み込まれます。DBの動作に関するオプションが設定されています。</b>
</ul>

<h3>model</h3>
<ul>
  <li><b>bookmark.go</b> - ブックマークのテーブルカラム定義、データの CRUD　処理が記述されています。</li>
  <li><b>user</b> - ユーザーデータのテーブルカラム定義、データの CRUD　処理が記述されています。</li>
</ul>

<h3>static</h3>
<ul>
  <li><b>css/style.css</b> - 共通CSSファイルです。</li>
  <li><b>images/</b> - テスト画像が格納されています。</li>
  <li><b>js/script.js</b> - 共通JSファイルです。</li>
</ul>

<h3>view</h3>
<ul>
  <li><b>base.html</b> - 共通ページのデザインが記述されています。</li>
  <li><b>bookmark.html</b> - ブックマークリストのデザインが記述されています。</li>
  <li><b>login.html</b> - ログインフォームのデザインが記述されています。</li>
  <li><b>menu.html</b> - メニューページのデザインが記述されています。</li>
  <li><b>menu.html</b> - メニュー(初期表示)ページのデザインが記述されています。</li>
  <li><b>password_reset.html</b> - パスワード再発行フォームのデザインが記述されています。</li>
  <li><b>profile.html</b> - プロフィールフォームのデザインが記述されています。</li>
  <li><b>search_books.html</b> - 書籍検索結果リストのデザインが記述されています。</li>
  <li><b>search_books_detail.html</b> - 書籍詳細情報のデザインが記述されています。</li>
  <li><b>signup.html</b> - ユーザー新規登録フォームのデザインが記述されています。</li>
</ul>

<h3>.env</h3>
<ul>
  <b>データベースの接続情報が記述されています。main.go から読み出され、DB 接続で使用されます。</b>
</ul>

<h3>Dockerfile</h3>
<ul>
  <b>Docker イメージの指定、ホットリロードツールである air のインストール、WORKDIR の変更が記述されています。</b>
</ul>

<h3>compose.yaml</h3>
<ul>
  <b>アプリサーバ(go)、DB サーバ(db)、DBクライアントツール(phpmyadmin)用のコンテナ定義が記述されています。</b>
</ul>

<h3>go.mod</h3>
<ul>
  <b>Go プログラミングで使われるモジュールの依存関係が記述されています。</b>
</ul>

<h3>go.sum</h3>
<ul>
  <b>go.mod ファイル内のモジュールの依存関係がハッシュ化された情報が記述されています。</b>
</ul>

<h3>main.go</h3>
<ul>
  <b>Go プログラムのエントリーポイントです。プログラム起動直後は、まず「main」パッケージに属する「main 関数」が実行されます。</b>
</ul>
