// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JacobBMay.Nutanix.Outputs
{

    [OutputType]
    public sealed class NdbClusterEntityCountEngineCount
    {
        public readonly ImmutableArray<Outputs.NdbClusterEntityCountEngineCountMariadbDatabase> MariadbDatabases;
        public readonly ImmutableArray<Outputs.NdbClusterEntityCountEngineCountMongodbDatabase> MongodbDatabases;
        public readonly ImmutableArray<Outputs.NdbClusterEntityCountEngineCountMysqlDatabase> MysqlDatabases;
        public readonly ImmutableArray<Outputs.NdbClusterEntityCountEngineCountOracleDatabase> OracleDatabases;
        public readonly ImmutableArray<Outputs.NdbClusterEntityCountEngineCountPostgresDatabase> PostgresDatabases;
        public readonly ImmutableArray<Outputs.NdbClusterEntityCountEngineCountSaphanaDatabase> SaphanaDatabases;
        public readonly ImmutableArray<Outputs.NdbClusterEntityCountEngineCountSqlserverDatabase> SqlserverDatabases;

        [OutputConstructor]
        private NdbClusterEntityCountEngineCount(
            ImmutableArray<Outputs.NdbClusterEntityCountEngineCountMariadbDatabase> mariadbDatabases,

            ImmutableArray<Outputs.NdbClusterEntityCountEngineCountMongodbDatabase> mongodbDatabases,

            ImmutableArray<Outputs.NdbClusterEntityCountEngineCountMysqlDatabase> mysqlDatabases,

            ImmutableArray<Outputs.NdbClusterEntityCountEngineCountOracleDatabase> oracleDatabases,

            ImmutableArray<Outputs.NdbClusterEntityCountEngineCountPostgresDatabase> postgresDatabases,

            ImmutableArray<Outputs.NdbClusterEntityCountEngineCountSaphanaDatabase> saphanaDatabases,

            ImmutableArray<Outputs.NdbClusterEntityCountEngineCountSqlserverDatabase> sqlserverDatabases)
        {
            MariadbDatabases = mariadbDatabases;
            MongodbDatabases = mongodbDatabases;
            MysqlDatabases = mysqlDatabases;
            OracleDatabases = oracleDatabases;
            PostgresDatabases = postgresDatabases;
            SaphanaDatabases = saphanaDatabases;
            SqlserverDatabases = sqlserverDatabases;
        }
    }
}