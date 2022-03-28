#! /bin/bash -e

#tag=v0.3.0
tag=$1
user=lfbdev
repo=stream
github_token=$GITHUB_TOKEN
plugin_dir=.devstream

#GOOS=$(go env GOOS)
#GOARCH=$(go env GOARCH)
GOOS=$2
GOARCH=$3

# upload dtm
echo 'Uploading 'dtm-${GOOS}-${GOARCH}' ...'
github-release upload --security-token $github_token --user $user --repo $repo --tag $tag --file dtm --name dtm-${GOOS}-${GOARCH}
echo dtm-${GOOS}-${GOARCH}' uploaded.'

# upload dtm .md5
echo 'Uploading 'dtm-${GOOS}-${GOARCH}.md5' ...'
github-release upload --security-token $github_token --user $user --repo $repo --tag $tag --file dtm-${GOOS}-${GOARCH}.md5 --name dtm-${GOOS}-${GOARCH}.md5
echo dtm-${GOOS}-${GOARCH}.md5' uploaded.'

# upload each plugin
for file in `ls $plugin_dir`
do
 if [ -d $plugin_dir"/"$file ]
 then
   read_dir $plugin_dir"/"$file
 else
   echo 'Uploading '$file' ...'
   github-release upload --security-token $github_token --user $user --repo $repo --tag $tag --file  $plugin_dir"/"$file --name $file
   echo $file' uploaded.'
 fi
done
