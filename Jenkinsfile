#!/usr/bin/env groovy

def config = [
        artifactId     : 'aurora-management-interface-go',
        groupId        : 'no.skatteetaten.aurora',
        scriptVersion  : 'v7',
        pipelineScript : 'https://git.aurora.skead.no/scm/ao/aurora-pipeline-scripts.git',
        goVersion      : 'Go 1.13',
        credentialsId  : "github",
        iq             : false,
        sonarQube      : false,
        versionStrategy: [
                [branch: 'master', versionHint: '0.1']
        ],
        debug          : true
]

run(config.scriptVersion, config)

def run(String scriptVersion, Map<String, Object> overrides = [:]) {

    if (overrides.debug) {
        println("Custom go build overrides: $overrides")
    }
    overrides.applicationType = "golanglib"

    def utilities
    def propertiesUtils
    def git
    def go

    Map props
    timestamps {
        node {

            stage('Load shared libraries') {
                fileLoader.withGit(overrides.pipelineScript, scriptVersion) {
                    utilities = fileLoader.load('utilities/utilities')
                    propertiesUtils = fileLoader.load('utilities/properties')
                    git = fileLoader.load('git/git')
                    go = fileLoader.load('go/go')

                }
            }

            stage('Checkout') {
                checkout scm
            }

            stage('Prepare') {
                props = propertiesUtils.getDefaultProps(overrides)
                utilities.initProps(props, git)

                if (overrides.debug) {
                    println("Custom go build - props: $props")
                }
            }

            stage('Test and coverage') {
                if (props.goVersion) {
                    go.buildGoWithJenkinsSh(props.goVersion)
                } else {
                    error("You need to specify goVersion")
                }
            }

            if (props.sonarQube) {
                stage('Sonar') {
                    def sonarPath = tool 'Sonar 4'
                    sh "${sonarPath}/bin/sonar-scanner -Dsonar.branch.name=${env.BRANCH_NAME}"
                }
            }

            if (props.isReleaseBuild && !props.tagExists) {
                stage("Tag") {
                    git.tagAndPush(props.credentialsId, "v$props.version")
                }
            }

            return props
        }
    }
}

return this