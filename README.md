##### Youtubeの歌枠
##### 推しの歌を登録
###### いつでも好きな歌を再生、布教にも。
###### 推しの歌を１画面で把握

～～アプリ開発中 & readme作成中～～


〇作成した目的

　youtubeでは１人の配信者が１配信で何曲も歌う「歌枠配信」がある。
 
しかし、どの動画で何をどの時間から歌ったかは視聴してみないと分からないことが多い。
この不便さ解消のため、当サイトへ動画urlと歌い出しの時間をメモし、いつでも好きな歌を聴けるwebサイトを作成した。

　登録した歌は表になり、一目で網羅できる
 
    サイトリンク：完成後に掲載予定

〇使用技術

- 言語、サービス


| 言語、サービス  | フレームワーク等 |
| ------------- | ------------- |
| JavaScript  | Next.js v13.4.12, TypeScript v5.1.6  |
| Go v1.18  | GORM, GIN  |
|MySQL v8.0.32| |
| Docker | |
| AWS EC | |
| Github | |


〇アーキテクチャ
 - REST API

〇技術選定理由

###### 業界未経験の私にとって以下の３点を目的に選定した。
　・学習しやすさ、参入しやすさ、入社後も長期活用できること

- フロントエンド
#### Next.js：広く普及しているReactベースのフレームワークであり、SSR, SSGを簡単に実装できること。
#### 基本的にTypeScriptであり、静的型付けでありGoと平衡で学びやすいこと。
 
-バックエンド
 
#### Go

    ・静的片付け言語であること

　　・厳格な記述ルールゆえの可読性の高さ
 
　　・コンパイル前のエラー検知
  
　　→この3つを業界未経験者にとって学習しやすいと考えた。

  
・求人数の多い言語の中で新生の言語

→廃れ流行りでは流行りに位置する。

　以上により、求職者数/求人数の比率が低く、参入ハードルが低いと考えました。
 

ER図
https://user-images.githubusercontent.com/127638412/273976430-29856108-a613-493e-b024-bb9ad7ac88d9.png

画面遷移図 main
https://user-images.githubusercontent.com/127638412/273976533-1d5db155-c5a6-403c-95db-f05b6fefb3f0.png

画面遷移図 header/footer
https://user-images.githubusercontent.com/127638412/273976737-17eb88c9-dfb0-4a99-a5ba-bf909c268ada.png

◯要件
--ver. 1--(デプロイし、自分で使用、バグ修正、微調整する)
〇機能要件
・ログイン機能（ゲストログイン有り）
・ユーザーは２種類：配信者とリスナー
・配信者に特別：厳正な会員登録(twitterと紐づけ？)、自身が関与するtableに対しては全権保持
・配信者、youtube歌枠動画、歌のデータ登録/編集/削除(会員専用)
・登録データを表で閲覧（非会員可能）
・youtubeでは1動画に2つのURLが有り、登録時に自動変換して片方へ統一
・削除、申請依頼
・いいね機能：配信者、歌に対して
・フォロー機能：リスナーは配信者やリスナーに対して
・バリデーション：メアド、パスワード、動画タイトル、その他色々

〇非機能要件　開発中、推敲中
・https化
・パスワードのbycript化
・N+1問題
・定期的な自動バックアップ
・ロック機能？
・死活監視
・名前ﾖﾐｶﾞﾅ、ﾊﾟｽﾜｰﾄﾞ登録時に全角を半角に自動変換
・メンテナンス中はサービス停止とする


--ver. 1.5--(知人にデバッグを兼ねて使用依頼、フィードバックを貰う)
〇機能要件
・スマホ対応デザイン
・ブラウザ対応：Chrome, Safari, FireFox, Opera, 


--ver. 2(Xで一般公開)--
〇機能要件
・退会機能
・パスワード変更機能

〇非機能要件
・ログイン認証のJWTの秘密鍵を定期的に自動更新


--ver. 3--(一般公開しながらのバージョンアップ)
〇機能要件
・オリ曲、カバー曲のtable追加
・"鳴き声"や"お話"のtabale追加
・お知らせ機能１：ログインした際に通知が届く。通知：フォロー関連の新規登録情報等
・会員登録時に本人確認メール
・同時視聴アプリ(youtubeのコメント欄ではリスナー同士の会話は忌まわれる。本機能では不明点等を教え合う場や、感想を言い合う場として利用できる。)
・配信者情報に項目追加：X, Fanbox, Booth, 画像
・曲登録時に±30s以内のものがある場合に警告する

〇非機能要件
・

--ver4
・X投稿機能
・お知らせ機能２：新規動画をメールで？初期は１通送って、通知onを促す
・開発者のXs表示(機能追加、デバッグ等のお知らせ)

--verX--
・ブラックリスト登録(IP BAN?)

◯マニュアル
###### 自分の好きな配信者とその歌情報（動画と歌い出しの時間等）を登録できます。
###### 歌情報については、登録情報の詳細ページで直接動画を再生することができます。

###### 会員機能
　　・歌情報登録
　　・自分が登録した歌情報の編集
    ※現状はゲストログイン機能を使用できます。
###### 非会員/会員共通機能
 　・歌情報の閲覧とページ内視聴

###### 主な用途
　・自分で登録した歌を見返すこと
　・布教に使う

