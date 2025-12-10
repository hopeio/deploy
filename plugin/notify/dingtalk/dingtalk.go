package dingtalk

import (
	"fmt"
	"time"

	"github.com/hopeio/gox/sdk/dingtalk"
)

type Config struct {
	Repo             string `flag:"name:repo;usage:deploy repo;env:DRONE_REPO"`
	Commit           string `flag:"name:commit;usage:commit;env:DRONE_COMMIT"`
	CommitTag        string `flag:"name:commit_tag;usage:commit_tag;env:DRONE_TAG"`
	CommitLink       string `flag:"name:commit_link;usage:commit_link;env:DRONE_COMMIT_LINK"`
	CommitRef        string `flag:"name:commit_ref;usage:commit_ref;env:DRONE_COMMIT_REF"`
	CommitMessage    string `flag:"name:commit_message;usage:commit_message;env:DRONE_COMMIT_MESSAGE"`
	CommitBranch     string `flag:"name:commit_branch;usage:commit_branch;env:DRONE_COMMIT_BRANCH"`
	CommitAuthor     string `flag:"name:commit_author;usage:git commit author;env:DRONE_COMMIT_AUTHOR"`
	CommitAuthorName string `flag:"name:commit_author_name;usage:git commit author name;env:DRONE_COMMIT_AUTHOR_NAME"`
	DingToken        string `flag:"name:ding_token;usage:ding_token;env:PLUGIN_DING_TOKEN"`
	DingSecret       string `flag:"name:ding_secret;usage:ding_secret;env:PLUGIN_DING_SECRET"`
	BuildLink        string `flag:"name:drone_build_link;usage:drone_build_link;env:DRONE_BUILD_LINK"`
}

func Notify(c *Config) error {

	if c.DingToken == "" {
		return nil
	}

	msg := "\\n # 发布通知 " +
		" \\n ### 项目: " + c.Repo +
		" \\n ### 操作人: " + c.CommitAuthor +
		" \\n ### 参考: " + c.CommitRef +
		" \\n ### 分支: " + c.CommitBranch +
		" \\n ### 标签: " + c.CommitTag +
		" \\n ### 时间: " + fmt.Sprint(time.Now().Format("2006-01-02 15:04:05")) +
		" \\n ### 提交: " + c.Commit +
		" \\n ### 提交信息: " + c.CommitMessage +
		" \\n ### 发布详情: " + c.BuildLink

	var err error
	if c.DingSecret != "" {
		err = dingtalk.RobotSendMarkDownMessageWithSecret(c.DingToken, c.DingSecret, "发布通知", msg, nil)
	} else {
		err = dingtalk.RobotSendMarkDownMessage(c.DingToken, "发布通知", msg, nil)
	}

	return err
}
