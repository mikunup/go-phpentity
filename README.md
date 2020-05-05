# go-phpentity

phpのセッター、ゲッターのクラスファイルを生成します。

# Usage

go run main.go ファイル名(クラス名) プロパティ名:型名:コメント

go run main.go User id:int name:string:名前

コメントは任意です。

# Sample

go run main.go User id:int name:string:名前

User.php

```

class User
{
    /** @var int */
    private $id;

    /** @var string 名前 */
    private $name;

    /**
     * Id Setter
     *
     * @param int id
     */
    public function setId($id)
    {
        $this->id = $id;
    }

    /**
     * Name Setter
     *
     * @param string name
     */
    public function setName($name)
    {
        $this->name = $name;
    }

    /**
     * Id Getter
     *
     * @return int id
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Name Getter
     *
     * @return string name
     */
    public function getName()
    {
        return $this->name;
    }

}

```

## Tool

For Linux mkpentity

For Windows mkpentity.exe
