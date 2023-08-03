ThisBuild / version := "0.1.0-SNAPSHOT"

ThisBuild / scalaVersion := "2.12.18"

lazy val root = (project in file("."))
  .settings(
    name := "consumer",
    idePackagePrefix := Some("org.xqdl")
  )

libraryDependencies ++= Seq(
  "org.apache.flink" % "flink-connector-kafka" % "1.16.1",
  "org.apache.flink" % "flink-clients" % "1.16.1",
  "org.apache.flink" %% "flink-streaming-scala" % "1.16.1" % "provided",
  "org.slf4j" % "slf4j-simple" % "2.0.5",
  "com.fasterxml.jackson.module" %% "jackson-module-scala" % "2.12.3"
)
