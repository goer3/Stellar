import { GithubOutlined } from '@ant-design/icons';

// 底部内容
const FooterDescriptionComponent = () => {
  // 运行环境
  const runEnv = window.CONFIG.env;
  const runEnvText = 'Running Env: ' + runEnv;
  const githubUrl = 'https://github.com/goer3/Stellar/releases';
  return (
    <>
      <b>👻 STELLAR </b>© 2024 EZOPS.CN. Latest Version:{' '}
      <a href={githubUrl} target="_blank" rel="noreferrer">
        <GithubOutlined />
      </a>{' '}
      / {runEnvText}
    </>
  );
};

// 项目描述
const ProjectDescription = 'Stellar 是一个集成了系统监控和业务监控，支持多数据源、多告警源、多告警通知方式、多告警处理方式的运维监控系统';

// 页面标题后缀
const TitleSuffix = ' | ' + ProjectDescription;

export { FooterDescriptionComponent, ProjectDescription, TitleSuffix };
